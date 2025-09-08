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

				fmt.Print(`
    __ __  _             __    
   / //_/ (_)____   ____/ /___ 
  / ,<   / // __ \ / __  // _ \
 / /| | / // / / // /_/ //  __/
/_/ |_|/_//_/ /_/ \__,_/ \___/      
`)
				fmt.Println()
				fmt.Printf("Version %v\n", release.Version)
				release.IsNeedingUpdate()
			},
		},
	}
}
