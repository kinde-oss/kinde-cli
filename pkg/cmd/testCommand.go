package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type testCmd struct {
	cmd *cobra.Command
}

func newTestCmd() *testCmd {
	return &testCmd{
		cmd: &cobra.Command{
			Use:   "test",
			Args:  nil,
			Short: "Runs tests",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("not implemented")
			},
		},
	}
}
