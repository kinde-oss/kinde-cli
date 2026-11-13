package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// envNoUpdateNotifier turns off the automatic update check.
const envNoUpdateNotifier = "KINDE_NO_UPDATE_NOTIFIER"

type versionCmd struct {
	cmd *cobra.Command
}

func newVersionCmd(ios *iostreams.IOStreams) *versionCmd {
	return &versionCmd{
		cmd: &cobra.Command{
			Use:   "version",
			Args:  cobra.NoArgs,
			Short: "Show the CLI's version and check for updates",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Fprintln(ios.Out, release.String())

				// The version is the output; the update check is a notice for
				// a person, so it's skipped when stderr isn't a terminal.
				if !ios.IsStderrTTY() {
					return nil
				}
				checker, err := newUpdateChecker()
				if err != nil {
					return nil
				}
				update, err := checker.Check(cmd.Context(), true)
				if err != nil {
					log.Ctx(cmd.Context()).Debug().Err(err).Msg("Update check failed")
				}
				if update != nil {
					fmt.Fprintf(ios.ErrOut, "\n%s\n", update.Message())
				}
				return nil
			},
		},
	}
}

func newUpdateChecker() (*release.Checker, error) {
	dir, err := config.Dir()
	if err != nil {
		return nil, err
	}
	return release.NewChecker(filepath.Join(dir, "update-check.json")), nil
}

// commandsWithoutUpdateCheck either check themselves (version) or must only
// print what the shell asked for (completion).
var commandsWithoutUpdateCheck = []string{"version", "completion", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd}

// startUpdateCheck checks for a newer release in the background while the
// command runs, at most once a day and only in a terminal. The channel
// yields the update, or nil.
func startUpdateCheck(ctx context.Context, ios *iostreams.IOStreams, args []string) <-chan *release.Update {
	result := make(chan *release.Update, 1)

	if !ios.IsStderrTTY() || os.Getenv(envNoUpdateNotifier) != "" || os.Getenv("CI") != "" ||
		(len(args) > 0 && slices.Contains(commandsWithoutUpdateCheck, args[0])) {
		close(result)
		return result
	}
	checker, err := newUpdateChecker()
	if err != nil {
		close(result)
		return result
	}

	go func() {
		defer close(result)
		update, err := checker.Check(ctx, false)
		if err != nil {
			log.Ctx(ctx).Debug().Err(err).Msg("Update check failed")
		}
		result <- update
	}()
	return result
}
