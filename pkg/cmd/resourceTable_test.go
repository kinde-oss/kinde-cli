package cmd

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/stretchr/testify/assert"
)

var kebabCase = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// deprecatedClientMethods are aliases kept by kinde-go for compatibility
// (see management_api/compat.go); the CLI uses the current names.
var deprecatedClientMethods = map[string]bool{
	"DeleteAPIAppliationScope": true, // use DeleteAPIApplicationScope
}

func TestResourceTable_CoversEveryOperationOnce(t *testing.T) {
	seen := map[management_api.OperationName]string{}
	for _, group := range resourceGroups {
		for _, command := range group.commands {
			path := group.name + " " + command.name
			if previous, ok := seen[command.operation]; ok {
				t.Errorf("operation %s is mapped twice: %q and %q", command.operation, previous, path)
			}
			seen[command.operation] = path
		}
	}

	client := reflect.TypeOf(&management_api.Client{})
	for i := 0; i < client.NumMethod(); i++ {
		name := client.Method(i).Name
		if _, ok := seen[name]; !ok && !deprecatedClientMethods[name] {
			t.Errorf("Management API operation %s has no CLI command; add it to resourceGroups", name)
		}
	}
}

func TestResourceTable_NamesAreUniqueKebabCase(t *testing.T) {
	groups := map[string]bool{}
	for _, group := range resourceGroups {
		assert.Regexp(t, kebabCase, group.name)
		assert.False(t, groups[group.name], "duplicate group %s", group.name)
		groups[group.name] = true

		commands := map[string]bool{}
		for _, command := range group.commands {
			assert.Regexp(t, kebabCase, command.name, group.name)
			assert.False(t, commands[command.name], "duplicate command %s %s", group.name, command.name)
			commands[command.name] = true
		}
	}
}
