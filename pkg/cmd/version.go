package cmd

import (
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	cmd *cobra.Command
}

func newVersionCmd() *versionCmd {
	return &versionCmd{
		cmd: &cobra.Command{
			Use:   "version",
			Args:  nil,
			Short: "Version",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("version %v\n", release.Branch)
				release.IsNeedingUpdate()
			},
		},
	}
}
