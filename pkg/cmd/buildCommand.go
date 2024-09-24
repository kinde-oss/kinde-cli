package cmd

import (
	"github.com/kinde-oss/kinde-cli/pkg/validators"
	builder "github.com/kinde-oss/kinde-cli/pkg/workflowBuilder"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type buildCmd struct {
	cmd *cobra.Command
	cwd string
}

func newBuildCmd() *buildCmd {
	dc := &buildCmd{}

	dc.cmd = &cobra.Command{
		Use:   "build",
		Args:  validators.NoArgs,
		Short: "Creating deployable version of code in ./kinde/output",
		Run: func(cmd *cobra.Command, args []string) {
			builder := builder.NewWorkflowBuilder(builder.BuildOptions{
				WorkingFolder: dc.cwd,
				EntryPoint:    "./*.ts",
			})
			result := builder.Build()
			log.Info().Msgf("Build result: %v", result)

		},
	}

	dc.cmd.Flags().StringVar(&dc.cwd, "cwd", "", "The directory to build from")

	return dc
}
