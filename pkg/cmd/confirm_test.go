package cmd

import (
	"bytes"
	"testing"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// answering returns a setup for runWith that makes the streams a terminal
// where the user types answer.
func answering(answer string) func(*iostreams.IOStreams, *bytes.Buffer) {
	return func(ios *iostreams.IOStreams, in *bytes.Buffer) {
		ios.SetTerminalInput(nil)
		in.WriteString(answer)
	}
}

func TestDestructiveCommand_NeedsYesOutsideATerminal(t *testing.T) {
	_, _, err := newLoggedInCLI(t).run("users", "delete", "--id", "kp_1")

	assert.EqualError(t, err, "'kinde users delete' can't be undone: pass --yes to confirm")
}

func TestDestructiveCommand_DecliningChangesNothing(t *testing.T) {
	stdout, stderr, err := newLoggedInCLI(t).runWith(answering("n\n"), "users", "delete", "--id", "kp_1")

	assert.EqualError(t, err, "cancelled: nothing was changed")
	assert.Equal(t, "Delete user --id kp_1 on "+testDomain+"? This can't be undone. [y/N] ", stderr)
	assert.Empty(t, stdout)
}

func TestDestructiveCommand_EmptyAnswerMeansNo(t *testing.T) {
	_, _, err := newLoggedInCLI(t).runWith(answering("\n"), "users", "delete", "--id", "kp_1")

	assert.EqualError(t, err, "cancelled: nothing was changed")
}

func TestDestructiveCommand_InvalidFlagsFailBeforeAsking(t *testing.T) {
	_, stderr, err := newLoggedInCLI(t).runWith(answering("y\n"), "users", "delete")

	assert.EqualError(t, err, "missing required flags: --id")
	assert.Empty(t, stderr, "there's nothing to confirm until the request is valid")
}

func TestConfirmDestructive(t *testing.T) {
	env := &config.Environment{DomainName: testDomain}
	op := resourceCommand{name: "delete", operation: management_api.DeleteUserOperation}

	newCmd := func(t *testing.T, args ...string) *cobra.Command {
		cmd, err := buildCobraCommand(op, nil)
		require.NoError(t, err)
		require.NoError(t, cmd.ParseFlags(args))
		return cmd
	}

	t.Run("yes answer", func(t *testing.T) {
		ios, in, _, _ := iostreams.Test()
		ios.SetTerminalInput(nil)
		in.WriteString("YES\n")
		assert.NoError(t, confirmDestructive(newCmd(t, "--id", "kp_1"), ios, env, op))
	})

	t.Run("--yes skips the question", func(t *testing.T) {
		ios, _, _, errOut := iostreams.Test()
		assert.NoError(t, confirmDestructive(newCmd(t, "--id", "kp_1", "--yes"), ios, env, op))
		assert.Empty(t, errOut.String())
	})

	t.Run("-y works too", func(t *testing.T) {
		ios, _, _, _ := iostreams.Test()
		assert.NoError(t, confirmDestructive(newCmd(t, "--id", "kp_1", "-y"), ios, env, op))
	})
}

func TestIsDestructive(t *testing.T) {
	assert.True(t, isDestructive(resourceCommand{operation: management_api.DeleteUserOperation}))
	assert.True(t, isDestructive(resourceCommand{operation: management_api.RemoveOrganizationUserOperation}), "DELETE operations whatever the name")
	assert.True(t, isDestructive(resourceCommand{operation: management_api.RotateApiKeyOperation}))
	assert.False(t, isDestructive(resourceCommand{operation: management_api.GetUserDataOperation}))
	assert.False(t, isDestructive(resourceCommand{operation: management_api.UpdateUserOperation}))
}

func TestDescribeFlags_TruncatesLongValues(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("body", "", "")
	cmd.Flags().Bool(yesFlag, false, "")
	long := string(bytes.Repeat([]byte("x"), maxConfirmValue+10))
	require.NoError(t, cmd.ParseFlags([]string{"--body", long, "--yes"}))

	assert.Equal(t, "--body "+long[:maxConfirmValue]+"...", describeFlags(cmd.Flags()))
}
