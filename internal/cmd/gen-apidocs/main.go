// Command gen-apidocs turns a Kinde API spec into Go source with the help
// text for each CLI command: the operation's summary, description, required
// scopes, its parameters, and a description for each parameter and body field.
//
// Usage: gen-apidocs <spec.yaml> <output.go> <variable name>
//
// It is run by go generate in pkg/cmd. To pick up API changes, replace the
// spec in api/ with the published one and run go generate ./...
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

// wrapWidth is the column the command descriptions are wrapped at.
const wrapWidth = 80

type operation struct {
	ID, Method, Path     string
	Summary, Description string
	Scopes               []string
	Params               []param
	Flags                map[string]string
}

// param is a path or query parameter.
type param struct {
	Name, In, Type string
	Required       bool
	Enum           []string
}

type spec struct {
	root map[string]any
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "usage: gen-apidocs <spec.yaml> <output.go> <variable name>")
		os.Exit(2)
	}
	if err := run(os.Args[1], os.Args[2], os.Args[3]); err != nil {
		fmt.Fprintln(os.Stderr, "gen-apidocs:", err)
		os.Exit(1)
	}
}

func run(specPath, outPath, varName string) error {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return err
	}
	var root map[string]any
	if err := yaml.Unmarshal(data, &root); err != nil {
		return fmt.Errorf("failed to parse %s: %w", specPath, err)
	}
	s := spec{root: root}

	operations, err := s.operations()
	if err != nil {
		return err
	}
	source, err := render(operations, filepath.Base(specPath), varName)
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, source, 0o644)
}

var methods = []string{"get", "put", "post", "delete", "options", "head", "patch", "trace"}

func (s spec) operations() ([]operation, error) {
	var operations []operation
	for path, item := range asMap(s.root["paths"]) {
		item := asMap(item)
		shared := asList(item["parameters"])

		for _, method := range methods {
			op := asMap(item[method])
			if op == nil {
				continue
			}
			id := asString(op["operationId"])
			if id == "" {
				return nil, fmt.Errorf("%s %s has no operationId", method, path)
			}

			description, scopes := cleanDescription(asString(op["description"]))
			o := operation{
				ID:          id,
				Method:      strings.ToUpper(method),
				Path:        path,
				Summary:     strings.TrimSuffix(strings.TrimSpace(asString(op["summary"])), "."),
				Description: description,
				Scopes:      scopes,
				Flags:       map[string]string{},
			}
			for _, param := range append(slices.Clone(shared), asList(op["parameters"])...) {
				param := s.resolve(param)
				o.addFlag(asString(param["name"]), asString(param["description"]))
				o.Params = append(o.Params, s.param(param))
			}
			s.addBodyFlags(&o, s.resolve(op["requestBody"]))
			operations = append(operations, o)
		}
	}
	slices.SortFunc(operations, func(a, b operation) int { return strings.Compare(a.ID, b.ID) })
	return operations, nil
}

func (s spec) param(p map[string]any) param {
	schema := s.resolve(p["schema"])
	var enum []string
	for _, value := range asList(schema["enum"]) {
		if value, ok := value.(string); ok {
			enum = append(enum, value)
		}
	}
	return param{
		Name:     asString(p["name"]),
		In:       asString(p["in"]),
		Type:     asString(schema["type"]),
		Required: p["required"] == true,
		Enum:     enum,
	}
}

func (s spec) addBodyFlags(o *operation, body map[string]any) {
	for _, mediaType := range []string{"application/json", "multipart/form-data", "application/x-www-form-urlencoded"} {
		content := asMap(asMap(body["content"])[mediaType])
		if content != nil {
			s.addSchemaFlags(o, content["schema"], "", 0)
			return
		}
	}
}

// addSchemaFlags records a description for each property of an object
// schema, recursing into nested objects with "parent.child" names to match
// the CLI's flags for nested fields.
func (s spec) addSchemaFlags(o *operation, schema any, prefix string, depth int) {
	resolved := s.resolve(schema)
	if resolved == nil || depth > 5 {
		return
	}
	for _, part := range asList(resolved["allOf"]) {
		s.addSchemaFlags(o, part, prefix, depth+1)
	}
	for name, property := range asMap(resolved["properties"]) {
		property := s.resolve(property)
		o.addFlag(prefix+name, asString(property["description"]))
		if asMap(property["properties"]) != nil || asList(property["allOf"]) != nil {
			s.addSchemaFlags(o, property, prefix+name+".", depth+1)
		}
	}
}

func (o *operation) addFlag(name, description string) {
	description = flagDescription(description)
	if name != "" && description != "" {
		o.Flags[docKey(name)] = description
	}
}

// resolve follows a local $ref ("#/components/schemas/X") and returns the
// referenced object.
func (s spec) resolve(value any) map[string]any {
	m := asMap(value)
	for range 10 {
		ref := asString(m["$ref"])
		if ref == "" {
			return m
		}
		var target any = s.root
		for _, part := range strings.Split(strings.TrimPrefix(ref, "#/"), "/") {
			target = asMap(target)[part]
		}
		m = asMap(target)
	}
	return m
}

// docKey normalizes a parameter or flag name so that "user_id", "user-id"
// and "UserID" all match. It must stay in sync with docKey in pkg/cmd.
func docKey(name string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '.' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

var (
	scopeBlock = regexp.MustCompile(`(?s)<div>\s*(.*?)\s*</div>`)
	scopeCode  = regexp.MustCompile(`<code>\s*([^<]+?)\s*</code>`)
	htmlTag    = regexp.MustCompile(`<[^>]+>`)
)

// cleanDescription removes the HTML block that lists the operation's required
// scopes, returning the scopes separately, and wraps the text for a terminal.
func cleanDescription(description string) (string, []string) {
	var scopes []string
	description = scopeBlock.ReplaceAllStringFunc(description, func(block string) string {
		for _, match := range scopeCode.FindAllStringSubmatch(block, -1) {
			scopes = append(scopes, match[1])
		}
		return ""
	})
	description = htmlTag.ReplaceAllString(description, "")

	var paragraphs []string
	for _, paragraph := range regexp.MustCompile(`\n\s*\n`).Split(strings.TrimSpace(description), -1) {
		if paragraph = strings.TrimSpace(paragraph); paragraph != "" {
			paragraphs = append(paragraphs, wrapParagraph(paragraph))
		}
	}
	return strings.Join(paragraphs, "\n\n"), scopes
}

var listItem = regexp.MustCompile(`^\s*([-*]|\d+\.)\s`)

// wrapParagraph wraps a paragraph to wrapWidth. Each list item starts on its
// own line, with its continuation lines indented.
func wrapParagraph(paragraph string) string {
	var items []string
	for _, line := range strings.Split(paragraph, "\n") {
		line = strings.TrimSpace(line)
		if len(items) == 0 || listItem.MatchString(line) {
			items = append(items, line)
		} else {
			items[len(items)-1] += " " + line
		}
	}
	for i, item := range items {
		indent := ""
		if listItem.MatchString(item) {
			indent = "  "
		}
		items[i] = wrap(item, indent)
	}
	return strings.Join(items, "\n")
}

func wrap(text, continuationIndent string) string {
	var b strings.Builder
	column := 0
	for i, word := range strings.Fields(text) {
		if i > 0 {
			if column+1+len(word) > wrapWidth {
				b.WriteString("\n" + continuationIndent)
				column = len(continuationIndent)
			} else {
				b.WriteByte(' ')
				column++
			}
		}
		b.WriteString(word)
		column += len(word)
	}
	return b.String()
}

// maxFlagDescription is the length above which a flag's description is cut to
// its first sentence, to keep each flag's help on one or two lines.
const maxFlagDescription = 120

// flagDescription makes a description fit on one line of flag help.
func flagDescription(description string) string {
	description = htmlTag.ReplaceAllString(description, "")
	description = strings.Join(strings.Fields(description), " ")
	// pflag treats the first `quoted` word as the flag's value name.
	description = strings.ReplaceAll(description, "`", "'")
	if len(description) > maxFlagDescription {
		description = firstSentence(description)
	}
	return strings.TrimSuffix(description, ".")
}

func firstSentence(text string) string {
	words := strings.Fields(text)
	for i, word := range words[:len(words)-1] {
		if strings.HasSuffix(word, ".") && !strings.Contains(strings.TrimSuffix(word, "."), ".") {
			return strings.Join(words[:i+1], " ")
		}
	}
	return text
}

func render(operations []operation, specName, varName string) ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by gen-apidocs from api/%s. DO NOT EDIT.\n\n", specName)
	b.WriteString("package cmd\n\n")
	fmt.Fprintf(&b, "var %s = map[string]operationDoc{\n", varName)
	for _, o := range operations {
		fmt.Fprintf(&b, "%q: {\n", strings.ToLower(o.ID))
		fmt.Fprintf(&b, "Method: %q,\nPath: %q,\n", o.Method, o.Path)
		fmt.Fprintf(&b, "Summary: %q,\n", o.Summary)
		fmt.Fprintf(&b, "Description: %q,\n", o.Description)
		if len(o.Scopes) > 0 {
			fmt.Fprintf(&b, "Scopes: %#v,\n", o.Scopes)
		}
		if len(o.Params) > 0 {
			b.WriteString("Params: []paramDoc{\n")
			for _, p := range o.Params {
				fmt.Fprintf(&b, "{Name: %q, In: %q, Type: %q", p.Name, p.In, p.Type)
				if p.Required {
					b.WriteString(", Required: true")
				}
				if len(p.Enum) > 0 {
					fmt.Fprintf(&b, ", Enum: %#v", p.Enum)
				}
				b.WriteString("},\n")
			}
			b.WriteString("},\n")
		}
		if len(o.Flags) > 0 {
			b.WriteString("Flags: map[string]string{\n")
			names := make([]string, 0, len(o.Flags))
			for name := range o.Flags {
				names = append(names, name)
			}
			slices.Sort(names)
			for _, name := range names {
				fmt.Fprintf(&b, "%q: %q,\n", name, o.Flags[name])
			}
			b.WriteString("},\n")
		}
		b.WriteString("},\n")
	}
	b.WriteString("}\n")
	return format.Source(b.Bytes())
}

func asMap(v any) map[string]any {
	m, _ := v.(map[string]any)
	return m
}

func asList(v any) []any {
	l, _ := v.([]any)
	return l
}

func asString(v any) string {
	s, _ := v.(string)
	return s
}
