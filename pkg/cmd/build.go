package cmd

import (
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/builder"
	"github.com/kinde-oss/kinde-cli/pkg/validators"
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
			builder := builder.NewBuilder()
			fmt.Printf("version %v - %v\n", builder, dc.cwd)
		},
	}

	dc.cmd.Flags().StringVar(&dc.cwd, "cwd", "", "The directory to build from")

	return dc
}
