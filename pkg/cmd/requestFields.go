package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/textproto"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ogen-go/ogen/http"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// fieldKind says how a flag's text becomes a request value.
type fieldKind int

const (
	kindString     fieldKind = iota // strings, including enums
	kindBool                        //
	kindInt                         //
	kindFloat                       //
	kindTime                        // RFC 3339 timestamp
	kindUUID                        //
	kindStringList                  // comma-separated or repeated flag
	kindFile                        // path to a file to upload
	kindJSON                        // anything else, given as JSON
)

const bodyFlag = "body"

var (
	timeType          = reflect.TypeOf(time.Time{})
	uuidType          = reflect.TypeOf(uuid.UUID{})
	multipartFileType = reflect.TypeOf(http.MultipartFile{})
)

type (
	// requestField is one value in an operation's arguments that a flag sets.
	requestField struct {
		flag     string
		arg      int   // index into requestPlan.args
		path     []int // struct field indexes from the argument to the value
		leaf     reflect.Type
		kind     fieldKind
		required bool
		enum     []string
	}

	// requestPlan describes how to build an operation's arguments from flags.
	// The same plan registers the flags and fills in the request, so the two
	// can't disagree.
	requestPlan struct {
		args    []reflect.Type // method arguments after the context
		bodyArg int            // index of the JSON request body in args, or -1
		fields  []requestField
	}
)

// newRequestPlan builds the plan for a Management API method's arguments
// (excluding the context).
func newRequestPlan(args []reflect.Type) (*requestPlan, error) {
	plan := &requestPlan{args: args, bodyArg: -1}

	for i, arg := range args {
		isParams := strings.HasSuffix(arg.Name(), "Params")
		inner, _ := unwrapType(arg)
		if !isParams && reflect.PointerTo(inner).Implements(jsonUnmarshalerType) {
			plan.bodyArg = i
		}
		plan.addFields(arg, i, "", nil, true)
	}

	seen := map[string]bool{bodyFlag: true}
	for _, field := range plan.fields {
		if seen[field.flag] {
			return nil, fmt.Errorf("duplicate flag --%s", field.flag)
		}
		seen[field.flag] = true
	}
	return plan, nil
}

var jsonUnmarshalerType = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()

func (p *requestPlan) addFields(t reflect.Type, arg int, flag string, path []int, required bool) {
	t, optional := unwrapType(t)
	required = required && !optional

	field := requestField{flag: flag, arg: arg, path: slices.Clone(path), leaf: t, required: required}

	switch {
	case t == timeType:
		field.kind = kindTime
	case t == uuidType:
		field.kind = kindUUID
	case t == multipartFileType:
		field.kind = kindFile
	case t.Kind() == reflect.String:
		field.kind = kindString
		field.enum = enumValues(t)
	case t.Kind() == reflect.Bool:
		field.kind = kindBool
	case t.Kind() >= reflect.Int && t.Kind() <= reflect.Uint64:
		field.kind = kindInt
	case t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64:
		field.kind = kindFloat
	case t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.String:
		field.kind = kindStringList
		field.enum = enumValues(t.Elem())
	case t.Kind() == reflect.Struct && !isOneOf(t):
		// A plain object: one flag per field, prefixed with this field's flag.
		for i := range t.NumField() {
			sf := t.Field(i)
			if sf.IsExported() {
				p.addFields(sf.Type, arg, flagNameFor(sf, flag), append(path, i), required)
			}
		}
		return
	default:
		field.kind = kindJSON
	}

	if flag != "" {
		p.fields = append(p.fields, field)
	}
}

// isWrapper reports whether t is one of ogen's optional or nullable wrappers
// (OptString, NilInt, OptNilUser...), which hold the real value in Value.
func isWrapper(t reflect.Type) bool {
	if t.Kind() != reflect.Struct || !(strings.HasPrefix(t.Name(), "Opt") || strings.HasPrefix(t.Name(), "Nil")) {
		return false
	}
	_, hasValue := t.FieldByName("Value")
	_, hasSet := t.FieldByName("Set")
	_, hasNull := t.FieldByName("Null")
	return hasValue && (hasSet || hasNull)
}

func unwrapPointer(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

// unwrapType returns the type a flag actually sets, looking through pointers
// and wrappers. optional is true if an Opt wrapper was crossed.
func unwrapType(t reflect.Type) (inner reflect.Type, optional bool) {
	for {
		t = unwrapPointer(t)
		if !isWrapper(t) {
			return t, optional
		}
		if _, ok := t.FieldByName("Set"); ok {
			optional = true
		}
		value, _ := t.FieldByName("Value")
		t = value.Type
	}
}

// unwrapValue is unwrapType for values: it allocates pointers and marks
// wrappers as set on the way down, and returns the settable inner value.
func unwrapValue(v reflect.Value) reflect.Value {
	for {
		switch {
		case v.Kind() == reflect.Pointer:
			if v.IsNil() {
				v.Set(reflect.New(v.Type().Elem()))
			}
			v = v.Elem()
		case isWrapper(v.Type()):
			if set := v.FieldByName("Set"); set.IsValid() {
				set.SetBool(true)
			}
			if null := v.FieldByName("Null"); null.IsValid() {
				null.SetBool(false)
			}
			v = v.FieldByName("Value")
		default:
			return v
		}
	}
}

// isOneOf reports whether t is an ogen sum type (a value that can take one of
// several shapes), which have IsXxx methods for each variant.
func isOneOf(t reflect.Type) bool {
	if _, ok := t.FieldByName("Type"); !ok {
		return false
	}
	for i := range t.NumMethod() {
		if strings.HasPrefix(t.Method(i).Name, "Is") {
			return true
		}
	}
	return false
}

// enumValues returns the allowed values of an ogen enum type, or nil.
func enumValues(t reflect.Type) []string {
	method, ok := t.MethodByName("AllValues")
	if !ok {
		return nil
	}
	values := method.Func.Call([]reflect.Value{reflect.Zero(t)})[0]
	enum := make([]string, values.Len())
	for i := range enum {
		enum[i] = values.Index(i).String()
	}
	return enum
}

// registerFlags adds a flag for every field, plus --body when the operation
// takes a JSON request body. describe returns a flag's description from the
// API docs, or "".
func (p *requestPlan) registerFlags(flags *pflag.FlagSet, describe func(flag string) string) {
	for _, field := range p.fields {
		usage := field.usage(describe(field.flag))
		switch field.kind {
		case kindBool:
			flags.Bool(field.flag, false, usage)
		case kindInt:
			flags.Int64(field.flag, 0, usage)
		case kindFloat:
			flags.Float64(field.flag, 0, usage)
		case kindStringList:
			flags.StringSlice(field.flag, nil, usage)
		default:
			flags.String(field.flag, "", usage)
		}
	}
	if p.bodyArg >= 0 {
		flags.String(bodyFlag, "", "Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields")
	}
}

// usage returns the flag's help: its description followed by how to give the
// value, e.g. "The identity type (one of: email, username; required)".
func (f requestField) usage(description string) string {
	var parts []string
	switch f.kind {
	case kindTime:
		parts = append(parts, "RFC 3339 timestamp, e.g. 2026-01-31T09:00:00Z")
	case kindFile:
		parts = append(parts, "path to a file")
	case kindJSON:
		parts = append(parts, "JSON value")
	case kindStringList:
		parts = append(parts, "comma-separated list")
	}
	if len(f.enum) > 0 {
		parts = append(parts, "one of: "+strings.Join(f.enum, ", "))
	}
	if f.required {
		parts = append(parts, "required")
	}
	hints := strings.Join(parts, "; ")
	switch {
	case description == "":
		return hints
	case hints == "":
		return description
	default:
		return description + " (" + hints + ")"
	}
}

// buildArgs creates the method arguments from the command's flags.
func (p *requestPlan) buildArgs(cmd *cobra.Command) ([]reflect.Value, error) {
	flags := cmd.Flags()

	args := make([]reflect.Value, len(p.args))
	for i, t := range p.args {
		if t.Kind() == reflect.Pointer {
			args[i] = reflect.New(t.Elem())
		} else {
			args[i] = reflect.New(t).Elem()
		}
	}

	bodyGiven := p.bodyArg >= 0 && flags.Changed(bodyFlag)
	if bodyGiven {
		if err := p.decodeBody(cmd, args[p.bodyArg]); err != nil {
			return nil, err
		}
	}

	var missing []string
	for _, field := range p.fields {
		if !flags.Changed(field.flag) {
			if field.required && !(bodyGiven && field.arg == p.bodyArg) {
				missing = append(missing, "--"+field.flag)
			}
			continue
		}
		value, err := field.parse(flags)
		if err != nil {
			return nil, fmt.Errorf("invalid value for --%s: %w", field.flag, err)
		}
		target := unwrapValue(args[field.arg])
		for _, i := range field.path {
			target = unwrapValue(target.Field(i))
		}
		target.Set(value)
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required flags: %s", strings.Join(missing, ", "))
	}

	return args, nil
}

func (p *requestPlan) decodeBody(cmd *cobra.Command, arg reflect.Value) error {
	source, _ := cmd.Flags().GetString(bodyFlag)

	var data []byte
	var err error
	switch {
	case source == "-":
		data, err = io.ReadAll(cmd.InOrStdin())
	case strings.HasPrefix(source, "@"):
		data, err = os.ReadFile(strings.TrimPrefix(source, "@"))
	default:
		data = []byte(source)
	}
	if err != nil {
		return fmt.Errorf("failed to read --%s: %w", bodyFlag, err)
	}

	target := unwrapValue(arg).Addr()
	if err := json.Unmarshal(data, target.Interface()); err != nil {
		return fmt.Errorf("invalid --%s: %w", bodyFlag, err)
	}
	return nil
}

// parse converts a flag's value to the field's Go type.
func (f requestField) parse(flags *pflag.FlagSet) (reflect.Value, error) {
	value := reflect.New(f.leaf).Elem()

	switch f.kind {
	case kindBool:
		b, err := flags.GetBool(f.flag)
		value.SetBool(b)
		return value, err
	case kindInt:
		n, err := flags.GetInt64(f.flag)
		if err != nil {
			return value, err
		}
		// Reject a value the field can't hold rather than send a different
		// number than the one typed.
		if value.CanUint() {
			if n < 0 || value.OverflowUint(uint64(n)) {
				return value, fmt.Errorf("%d is out of range", n)
			}
			value.SetUint(uint64(n))
		} else {
			if value.OverflowInt(n) {
				return value, fmt.Errorf("%d is out of range", n)
			}
			value.SetInt(n)
		}
		return value, nil
	case kindFloat:
		n, err := flags.GetFloat64(f.flag)
		value.SetFloat(n)
		return value, err
	case kindStringList:
		items, err := flags.GetStringSlice(f.flag)
		if err != nil {
			return value, err
		}
		list := reflect.MakeSlice(f.leaf, len(items), len(items))
		for i, item := range items {
			if err := f.checkEnum(item); err != nil {
				return value, err
			}
			list.Index(i).SetString(item)
		}
		return list, nil
	}

	s, err := flags.GetString(f.flag)
	if err != nil {
		return value, err
	}

	switch f.kind {
	case kindString:
		if err := f.checkEnum(s); err != nil {
			return value, err
		}
		value.SetString(s)
	case kindTime:
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return value, fmt.Errorf("expected an RFC 3339 timestamp such as 2026-01-31T09:00:00Z")
		}
		value.Set(reflect.ValueOf(t))
	case kindUUID:
		id, err := uuid.Parse(s)
		if err != nil {
			return value, err
		}
		value.Set(reflect.ValueOf(id))
	case kindFile:
		file, err := readMultipartFile(s)
		if err != nil {
			return value, err
		}
		value.Set(reflect.ValueOf(file))
	case kindJSON:
		data := []byte(s)
		// Allow plain text where a JSON string is expected: --value hello.
		if !json.Valid(data) {
			data = []byte(strconv.Quote(s))
		}
		if err := json.Unmarshal(data, value.Addr().Interface()); err != nil {
			return value, err
		}
	}
	return value, nil
}

func (f requestField) checkEnum(s string) error {
	if len(f.enum) > 0 && !slices.Contains(f.enum, s) {
		return fmt.Errorf("%q is not one of: %s", s, strings.Join(f.enum, ", "))
	}
	return nil
}

func readMultipartFile(path string) (http.MultipartFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return http.MultipartFile{}, err
	}
	name := filepath.Base(path)
	contentType := mime.TypeByExtension(filepath.Ext(name))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	header := textproto.MIMEHeader{}
	header.Set("Content-Type", contentType)
	return http.MultipartFile{Name: name, Header: header, File: bytes.NewReader(data)}, nil
}
