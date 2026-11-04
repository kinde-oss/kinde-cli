package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const yesFlag = "yes"

// irreversibleOperations can't be undone even though they aren't HTTP
// DELETEs. Every DELETE operation is confirmed too.
var irreversibleOperations = map[management_api.OperationName]bool{
	management_api.RotateApiKeyOperation:            true, // the old key stops working
	management_api.RevokeConnectedAppTokenOperation: true,
}

func isDestructive(op resourceCommand) bool {
	return docsFor(op.operation).Method == http.MethodDelete || irreversibleOperations[op.operation]
}

// confirmDestructive asks before running a command that can't be undone.
// Without a terminal to ask in, --yes is required, so a script can't delete
// something by accident.
func confirmDestructive(cmd *cobra.Command, ios *iostreams.IOStreams, env *config.Environment, op resourceCommand) error {
	if yes, _ := cmd.Flags().GetBool(yesFlag); yes {
		return nil
	}
	if !ios.CanPrompt() {
		return fmt.Errorf("'%s' can't be undone: pass --%s to confirm", cmd.CommandPath(), yesFlag)
	}

	question := docsFor(op.operation).Summary
	if target := describeFlags(cmd.Flags()); target != "" {
		question += " " + target
	}
	question += fmt.Sprintf(" on %s? This can't be undone.", env.DomainName)

	confirmed, err := ios.Confirm(question)
	if err != nil {
		return err
	}
	if !confirmed {
		return errors.New("cancelled: nothing was changed")
	}
	return nil
}

// maxConfirmValue is how much of each flag value the confirmation shows.
const maxConfirmValue = 40

// describeFlags lists the flags given on the command line, e.g.
// "--user-id kp_123", so the prompt shows what will be affected.
func describeFlags(flags *pflag.FlagSet) string {
	var parts []string
	flags.Visit(func(f *pflag.Flag) {
		if f.Name == yesFlag || f.Name == "domain" {
			return
		}
		value := f.Value.String()
		if len(value) > maxConfirmValue {
			value = value[:maxConfirmValue] + "..."
		}
		parts = append(parts, "--"+f.Name+" "+value)
	})
	return strings.Join(parts, " ")
}
