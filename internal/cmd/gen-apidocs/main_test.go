package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanDescription_ExtractsScopes(t *testing.T) {
	description, scopes := cleanDescription("Creates an identity.\n\n<div>\n  <code>create:user_identities</code>\n</div>\n")
	assert.Equal(t, "Creates an identity.", description)
	assert.Equal(t, []string{"create:user_identities"}, scopes)
}

func TestCleanDescription_WrapsParagraphsAndKeepsLists(t *testing.T) {
	long := strings.Repeat("word ", 30)
	description, _ := cleanDescription(long + "\n\n- first item\n- second item\n")

	for _, line := range strings.Split(description, "\n") {
		assert.LessOrEqual(t, len(line), wrapWidth, line)
	}
	assert.True(t, strings.HasSuffix(description, "\n\n- first item\n- second item"), description)
}

func TestCleanDescription_JoinsWrappedListItems(t *testing.T) {
	description, _ := cleanDescription("Rules:\n\n- `absolute`: replaces the total with\n  the value.\n- `delta`: adds the value.")
	assert.Equal(t, "Rules:\n\n- `absolute`: replaces the total with the value.\n- `delta`: adds the value.", description)
}

func TestFlagDescription(t *testing.T) {
	assert.Equal(t, "The user's ID", flagDescription("The user's ID."))
	// pflag would turn `value` into the flag's value name.
	assert.Equal(t, "Use 'ALL' for every type", flagDescription("Use `ALL` for\n every type."))

	long := "The identity value, which depends on the type. " + strings.Repeat("More detail, e.g. this. ", 10)
	assert.Equal(t, "The identity value, which depends on the type", flagDescription(long))
}

func TestDocKey(t *testing.T) {
	assert.Equal(t, "userid", docKey("user_id"))
	assert.Equal(t, "userid", docKey("UserID"))
	assert.Equal(t, "options.clientid", docKey("options.client_id"))
}

func TestResolve_FollowsRefs(t *testing.T) {
	s := spec{root: map[string]any{
		"components": map[string]any{"parameters": map[string]any{
			"user_id": map[string]any{"name": "user_id", "description": "The user's ID."},
		}},
	}}
	param := s.resolve(map[string]any{"$ref": "#/components/parameters/user_id"})
	assert.Equal(t, "user_id", param["name"])
}
