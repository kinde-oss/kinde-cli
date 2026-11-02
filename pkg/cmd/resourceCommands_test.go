package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFlagNameFor(t *testing.T) {
	type sample struct {
		Tagged   string `json:"phone_country_id"`
		Untagged string
		Option   string `json:"name,omitempty"`
	}
	typ := reflect.TypeOf(sample{})

	assert.Equal(t, "phone-country-id", flagNameFor(typ.Field(0), ""))
	assert.Equal(t, "untagged", flagNameFor(typ.Field(1), ""))
	assert.Equal(t, "name", flagNameFor(typ.Field(2), ""))
	assert.Equal(t, "colors.phone-country-id", flagNameFor(typ.Field(0), "colors"))
}

func TestEveryResourceCommandBuilds(t *testing.T) {
	for _, group := range resourceGroups {
		for _, op := range group.commands {
			_, err := buildCobraCommand(op)
			assert.NoError(t, err, "%s %s", group.name, op.name)
		}
	}
}
