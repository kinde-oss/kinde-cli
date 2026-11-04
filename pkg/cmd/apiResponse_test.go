package cmd

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// callCreateIdentity runs create-identity against a fake API that answers
// with the given status, headers and body.
func callCreateIdentity(t *testing.T, status int, header http.Header, body string) (*http.Request, string, error) {
	t.Helper()
	var received *http.Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = r
		for k, v := range header {
			w.Header()[k] = v
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	args, err := buildTestArgs(t, management_api.CreateUserIdentityOperation, "",
		"--user-id", "kp_1", "--type", "email", "--value", "sally@example.com")
	require.NoError(t, err)

	var out bytes.Buffer
	cmd := &cobra.Command{}
	cmd.SetOut(&out)
	cmd.SetContext(context.Background())
	token := func(context.Context) (string, error) { return "test-token", nil }

	err = callOperation(cmd, server.URL, server.Client(), token, management_api.CreateUserIdentityOperation, args)
	return received, out.String(), err
}

func TestCallOperation_PrintsResponseAsJSON(t *testing.T) {
	req, out, err := callCreateIdentity(t, http.StatusCreated, nil,
		`{"message":"Success","code":"IDENTITY_CREATED","identity":{"id":"identity_1"},"added_later":true}`)
	require.NoError(t, err)

	assert.Equal(t, "/api/v1/users/kp_1/identities", req.URL.Path)
	assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
	assert.Equal(t, "kinde-cli/dev", req.Header.Get("User-Agent"))

	// The body is printed as sent, including fields the SDK doesn't know.
	assert.Equal(t, `{
  "message": "Success",
  "code": "IDENTITY_CREATED",
  "identity": {
    "id": "identity_1"
  },
  "added_later": true
}
`, out)
}

func TestCallOperation_SuccessTheSDKCannotDecodeStillSucceeds(t *testing.T) {
	_, out, err := callCreateIdentity(t, http.StatusCreated, nil, `{"identity":{"id":12345}}`)
	require.NoError(t, err)
	assert.Contains(t, out, `"id": 12345`)
}

func TestCallOperation_ForbiddenNamesTheRequiredScope(t *testing.T) {
	_, out, err := callCreateIdentity(t, http.StatusForbidden, nil, "")
	assert.EqualError(t, err, "the API returned 403 Forbidden. Check that your M2M application has the create:user_identities scope for the Kinde Management API")
	assert.Empty(t, out)
}

func TestCallOperation_BadRequestShowsKindeErrors(t *testing.T) {
	_, _, err := callCreateIdentity(t, http.StatusBadRequest, nil,
		`{"errors":[{"code":"INVALID_EMAIL","message":"Email is invalid"},{"code":"OTHER","message":"Also wrong"}]}`)
	assert.EqualError(t, err, "the API returned 400 Bad Request: INVALID_EMAIL: Email is invalid; OTHER: Also wrong")
}

func TestCallOperation_TooManyRequestsSaysWhenToRetry(t *testing.T) {
	_, _, err := callCreateIdentity(t, http.StatusTooManyRequests, http.Header{"Retry-After": {"30"}}, "")
	assert.EqualError(t, err, "the API returned 429 Too Many Requests. Try again in 30 seconds")
}

func TestCallOperation_UnauthorizedSuggestsLoggingInAgain(t *testing.T) {
	_, _, err := callCreateIdentity(t, http.StatusUnauthorized, nil, `{"error":"invalid_token"}`)
	assert.ErrorContains(t, err, "401 Unauthorized: invalid_token")
	assert.ErrorContains(t, err, "run 'kinde login' again")
}

func TestErrorMessages_NonJSONBodyIsShownTruncated(t *testing.T) {
	assert.Equal(t, []string{"upstream timeout"}, errorMessages([]byte("upstream timeout\n")))
	long := bytes.Repeat([]byte("x"), maxRawErrorBody+50)
	assert.Len(t, errorMessages(long)[0], maxRawErrorBody+len("..."))
}

func TestWriteJSON_EmptyBodyPrintsNothing(t *testing.T) {
	var out bytes.Buffer
	require.NoError(t, writeJSON(&out, nil))
	assert.Empty(t, out.String())
}
