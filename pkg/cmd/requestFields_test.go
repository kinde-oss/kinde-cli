package cmd

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// buildTestArgs parses flags for an operation and builds its request
// arguments, without calling the API.
func buildTestArgs(t *testing.T, operation management_api.OperationName, stdin string, flags ...string) ([]reflect.Value, error) {
	t.Helper()
	plan, err := planForOperation(operation)
	require.NoError(t, err)
	cmd := &cobra.Command{Use: "test"}
	cmd.SetIn(strings.NewReader(stdin))
	plan.registerFlags(cmd.Flags(), docsFor(operation).flag)
	require.NoError(t, cmd.ParseFlags(flags))
	return plan.buildArgs(cmd)
}

// Issue #52: --type was missing because optional enum fields were skipped.
func TestCreateUserIdentity_SetsTypeAndValue(t *testing.T) {
	args, err := buildTestArgs(t, management_api.CreateUserIdentityOperation, "",
		"--user-id", "kp_1", "--type", "email", "--value", "sally@example.com")
	require.NoError(t, err)

	req := args[0].Interface().(management_api.OptCreateUserIdentityReq)
	require.True(t, req.Set)
	assert.Equal(t, management_api.NewOptCreateUserIdentityReqType("email"), req.Value.Type)
	assert.Equal(t, management_api.NewOptString("sally@example.com"), req.Value.Value)
	assert.False(t, req.Value.PhoneCountryID.Set, "unset flags must stay unset")

	params := args[1].Interface().(management_api.CreateUserIdentityParams)
	assert.Equal(t, "kp_1", params.UserID)
}

func TestEnumFlag_RejectsUnknownValue(t *testing.T) {
	_, err := buildTestArgs(t, management_api.CreateUserIdentityOperation, "",
		"--user-id", "kp_1", "--type", "fax")
	assert.ErrorContains(t, err, `invalid value for --type: "fax" is not one of`)
	assert.ErrorContains(t, err, "email")
}

func TestEnumFlag_ListsAllowedValuesInHelp(t *testing.T) {
	plan, err := planForOperation(management_api.CreateUserIdentityOperation)
	require.NoError(t, err)
	cmd := &cobra.Command{}
	plan.registerFlags(cmd.Flags(), docsFor(management_api.CreateUserIdentityOperation).flag)
	assert.Contains(t, cmd.Flags().Lookup("type").Usage, "one of: ")
}

func TestRequiredFlags_AreEnforced(t *testing.T) {
	_, err := buildTestArgs(t, management_api.CreateUserIdentityOperation, "", "--type", "email")
	assert.EqualError(t, err, "missing required flags: --user-id")
}

func TestBody_FlagsOverrideBodyFields(t *testing.T) {
	args, err := buildTestArgs(t, management_api.CreateUserIdentityOperation, "",
		"--user-id", "kp_1",
		"--body", `{"type":"email","value":"from-body@example.com"}`,
		"--value", "from-flag@example.com")
	require.NoError(t, err)

	req := args[0].Interface().(management_api.OptCreateUserIdentityReq)
	assert.Equal(t, management_api.NewOptCreateUserIdentityReqType("email"), req.Value.Type)
	assert.Equal(t, management_api.NewOptString("from-flag@example.com"), req.Value.Value)
}

func TestBody_FromStdin(t *testing.T) {
	args, err := buildTestArgs(t, management_api.CreateUserIdentityOperation,
		`{"type":"username","value":"sally"}`, "--user-id", "kp_1", "--body", "-")
	require.NoError(t, err)

	req := args[0].Interface().(management_api.OptCreateUserIdentityReq)
	assert.Equal(t, management_api.NewOptString("sally"), req.Value.Value)
}

func TestBody_SatisfiesRequiredBodyFields(t *testing.T) {
	keyType := management_api.CreateApiKeyReqType("").AllValues()[0]
	_, err := buildTestArgs(t, management_api.CreateApiKeyOperation, "",
		"--body", `{"name":"ci","type":"`+string(keyType)+`","api_id":"api_1"}`)
	assert.NoError(t, err)
}

func TestBody_InvalidJSON(t *testing.T) {
	_, err := buildTestArgs(t, management_api.CreateUserIdentityOperation, "",
		"--user-id", "kp_1", "--body", `{"type":`)
	assert.ErrorContains(t, err, "invalid --body")
}

func TestJSONFlag_OneOfAcceptsPlainTextOrJSON(t *testing.T) {
	build := func(value string) management_api.UpdateApplicationsPropertyReqValue {
		args, err := buildTestArgs(t, management_api.UpdateApplicationsPropertyOperation, "",
			"--application-id", "app_1", "--property-key", "plan", "--value", value)
		require.NoError(t, err)
		return args[0].Interface().(*management_api.UpdateApplicationsPropertyReq).Value
	}

	assert.Equal(t, management_api.NewStringUpdateApplicationsPropertyReqValue("pro"), build("pro"))
	assert.Equal(t, management_api.NewBoolUpdateApplicationsPropertyReqValue(true), build("true"))
}

func TestJSONFlag_ListOfObjects(t *testing.T) {
	args, err := buildTestArgs(t, management_api.AddOrganizationUsersOperation, "",
		"--org-code", "org_1", "--users", `[{"id":"kp_1"},{"id":"kp_2"}]`)
	require.NoError(t, err)

	req := args[0].Interface().(management_api.OptAddOrganizationUsersReq)
	require.Len(t, req.Value.Users, 2)
	assert.Equal(t, management_api.NewOptString("kp_2"), req.Value.Users[1].ID)
}

func TestStringListFlag(t *testing.T) {
	keyType := management_api.CreateApiKeyReqType("").AllValues()[0]
	args, err := buildTestArgs(t, management_api.CreateApiKeyOperation, "",
		"--name", "ci", "--type", string(keyType), "--api-id", "api_1", "--scope-ids", "s1,s2")
	require.NoError(t, err)

	req := args[0].Interface().(*management_api.CreateApiKeyReq)
	assert.Equal(t, []string{"s1", "s2"}, req.ScopeIds.Value)
	assert.True(t, req.ScopeIds.Set)
}

func TestTimeFlag(t *testing.T) {
	base := []string{"--customer-agreement-id", "a", "--billing-feature-code", "f", "--meter-value", "1"}

	args, err := buildTestArgs(t, management_api.CreateMeterUsageRecordOperation, "",
		append(base, "--meter-usage-timestamp", "2026-01-31T09:00:00Z")...)
	require.NoError(t, err)
	req := args[0].Interface().(*management_api.CreateMeterUsageRecordReq)
	assert.Equal(t, time.Date(2026, 1, 31, 9, 0, 0, 0, time.UTC), req.MeterUsageTimestamp.Value)

	_, err = buildTestArgs(t, management_api.CreateMeterUsageRecordOperation, "",
		append(base, "--meter-usage-timestamp", "yesterday")...)
	assert.ErrorContains(t, err, "RFC 3339")
}

func TestInvalidFlags_ReportedBeforeMissingDomain(t *testing.T) {
	_, _, err := newTestCLI(t).run("users", "create-identity", "--user-id", "kp_1", "--type", "fax")
	assert.ErrorContains(t, err, `invalid value for --type: "fax"`)
}
