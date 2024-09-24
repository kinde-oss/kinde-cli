package validators

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// validating that command doesn't support arguments
func NoArgs(cmd *cobra.Command, args []string) error {
	commandPath := cmd.CommandPath()
	errorMessage := fmt.Sprintf(
		"`%s` does not take any positional arguments. See `%s --help` for supported flags and usage",
		commandPath,
		commandPath,
	)

	if len(args) > 0 {
		return errors.New(errorMessage)
	}

	return nil
}
