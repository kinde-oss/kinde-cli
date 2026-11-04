package cmd

import (
	"strings"
)

//go:generate go run ../../internal/cmd/gen-apidocs ../../api/kinde-management-api-spec.yaml apiDocs_gen.go
//go:generate go run ../../internal/cmd/gen-docs ../../docs/commands

// operationDoc is the help text for a Management API operation, generated
// from the API spec into apiDocs_gen.go.
type operationDoc struct {
	Method, Path string
	Summary      string
	Description  string
	Scopes       []string
	// Flags maps docKey(flag name) to the parameter or body field's description.
	Flags map[string]string
}

// docsFor returns the help text for an operation. Operations missing from the
// spec get empty docs.
func docsFor(operation string) operationDoc {
	return apiDocs[strings.ToLower(operation)]
}

// flag returns the description for a flag, or "".
func (d operationDoc) flag(name string) string {
	return d.Flags[docKey(name)]
}

// long returns the command's long help: the description, the scopes the M2M
// application needs, and the API endpoint it calls.
func (d operationDoc) long() string {
	var sections []string
	if d.Description != "" {
		sections = append(sections, d.Description)
	}
	if len(d.Scopes) > 0 {
		label := "Required scope: "
		if len(d.Scopes) > 1 {
			label = "Required scopes: "
		}
		sections = append(sections, label+strings.Join(d.Scopes, ", "))
	}
	if d.Path != "" {
		sections = append(sections, "API: "+d.Method+" "+d.Path)
	}
	return strings.Join(sections, "\n\n")
}

// docKey normalizes a parameter or flag name so that "user_id", "user-id"
// and "UserID" all match. It must stay in sync with docKey in gen-apidocs.
func docKey(name string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '.' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
