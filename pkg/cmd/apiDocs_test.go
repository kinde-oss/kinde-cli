package cmd

import (
	"testing"

	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// If this fails after upgrading kinde-go, the pinned spec is older than the
// SDK: update api/kinde-management-api-spec.yaml and run go generate ./...
func TestEveryResourceCommandHasAPIDocs(t *testing.T) {
	for _, group := range resourceGroups {
		for _, op := range group.commands {
			doc := docsFor(op.operation)
			assert.NotEmpty(t, doc.Summary, "%s %s (%s) has no summary", group.name, op.name, op.operation)
			assert.NotEmpty(t, doc.Path, "%s %s (%s) has no API path", group.name, op.name, op.operation)
		}
	}
}

func TestResourceCommand_HelpComesFromSpec(t *testing.T) {
	command, err := buildCobraCommand(resourceCommand{name: "create-identity", operation: management_api.CreateUserIdentityOperation}, nil)
	require.NoError(t, err)

	assert.Equal(t, "Create identity", command.Short)
	assert.Contains(t, command.Long, "Creates an identity for a user.")
	assert.Contains(t, command.Long, "Required scope: create:user_identities")
	assert.Contains(t, command.Long, "API: POST /api/v1/users/{user_id}/identities")
	assert.NotContains(t, command.Long, "<code>")

	assert.Equal(t, "The user's ID (required)", command.Flags().Lookup("user-id").Usage)
	assert.Equal(t, "The identity type (one of: email, username, phone, enterprise, social)",
		command.Flags().Lookup("type").Usage)
}

func TestDocKey_MatchesFlagAndSpecNames(t *testing.T) {
	assert.Equal(t, docKey("user_id"), docKey("user-id"))
	assert.Equal(t, docKey("user_id"), docKey("UserID"))
	assert.Equal(t, "options.clientid", docKey("options.client-id"))
}

func TestOperationDoc_Long(t *testing.T) {
	assert.Equal(t, "", operationDoc{}.long())
	assert.Equal(t, "Does things.\n\nRequired scopes: read:a, read:b\n\nAPI: GET /x",
		operationDoc{Description: "Does things.", Scopes: []string{"read:a", "read:b"}, Method: "GET", Path: "/x"}.long())
}
