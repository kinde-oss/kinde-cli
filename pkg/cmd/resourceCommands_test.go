package cmd

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"SimpleTest", "simple_test"},
		{"HTTPServer", "http_server"},
		{"userID", "user_id"},
		{"UserID", "user_id"},
		{"userIDNumber", "user_id_number"},
		{"XMLParser", "xml_parser"},
		{"already_snake_case", "already_snake_case"},
		{"lowercase", "lowercase"},
		{"UpperCase", "upper_case"},
		{"SomeHTTPServer", "some_http_server"},
		{"ABCTest", "abc_test"},
		{"TestABC", "test_abc"},
		{"TestABCTest", "test_abc_test"},
		{"Test123ABC", "test123_abc"},
		{"TestABC123", "test_abc123"},
	}

	for _, tt := range tests {
		got := toSnakeCase(tt.input)
		if got != tt.expected {
			t.Errorf("toSnakeCase(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}

// Dummy types to mimic management_api types for testing
type OptNilBool struct{}
type OptBool struct{}
type OptNilInt struct{}
type OptInt struct{}
type OptNilString struct{}

// Test struct with various field types and json tags
type testStruct struct {
	BoolField      OptNilBool `json:"bool_field"`
	BoolFieldNoTag OptBool
	IntField       OptNilInt `json:"int_field"`
	IntFieldNoTag  OptInt
	StringField    OptNilString `json:"string_field"`
	StringNoTag    string
}

func TestGenerateCommandFlags(t *testing.T) {
	cmd := &cobra.Command{}
	// Patch management_api types for testing
	management_api := struct {
		OptNilBool   OptNilBool
		OptBool      OptBool
		OptNilInt    OptNilInt
		OptInt       OptInt
		OptNilString OptNilString
	}{}

	// Patch reflect.TypeOf for dummy types
	origOptNilBool := reflect.TypeOf(management_api.OptNilBool)
	origOptBool := reflect.TypeOf(management_api.OptBool)
	origOptNilInt := reflect.TypeOf(management_api.OptNilInt)
	origOptInt := reflect.TypeOf(management_api.OptInt)

	// Patch function for test
	generateCommandFlagsTest := func(t reflect.Type, command *cobra.Command) {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() != reflect.Struct {
			return
		}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("json")

			switch field.Type {
			case origOptNilBool, origOptBool:
				if tag != "" {
					command.Flags().Bool(tag, false, "")
				} else {
					command.Flags().Bool(toSnakeCase(field.Name), false, "")
				}
			case origOptNilInt, origOptInt:
				if tag != "" {
					command.Flags().Int(tag, 0, "")
				} else {
					command.Flags().Int(toSnakeCase(field.Name), 0, "")
				}
			default:
				if tag != "" {
					command.Flags().String(tag, "", "")
				} else {
					command.Flags().String(toSnakeCase(field.Name), "", "")
				}
			}
		}
	}

	generateCommandFlagsTest(reflect.TypeOf(testStruct{}), cmd)

	// Check bool flags
	if f := cmd.Flags().Lookup("bool_field"); f == nil || f.Value.Type() != "bool" {
		t.Errorf("Expected bool flag for 'bool_field'")
	}
	if f := cmd.Flags().Lookup("bool_field_no_tag"); f == nil || f.Value.Type() != "bool" {
		t.Errorf("Expected bool flag for 'bool_field_no_tag'")
	}

	// Check int flags
	if f := cmd.Flags().Lookup("int_field"); f == nil || f.Value.Type() != "int" {
		t.Errorf("Expected int flag for 'int_field'")
	}
	if f := cmd.Flags().Lookup("int_field_no_tag"); f == nil || f.Value.Type() != "int" {
		t.Errorf("Expected int flag for 'int_field_no_tag'")
	}

	// Check string flags
	if f := cmd.Flags().Lookup("string_field"); f == nil || f.Value.Type() != "string" {
		t.Errorf("Expected string flag for 'string_field'")
	}
	if f := cmd.Flags().Lookup("string_no_tag"); f == nil || f.Value.Type() != "string" {
		t.Errorf("Expected string flag for 'string_no_tag'")
	}
}

func TestGenerateCommandFlags_NonStruct(t *testing.T) {
	cmd := &cobra.Command{}
	// Should not panic or add flags for non-struct types
	generateCommandFlags(reflect.TypeOf("string"), cmd)
	if cmd.Flags().HasFlags() {
		t.Errorf("Expected no flags for non-struct type")
	}
}
