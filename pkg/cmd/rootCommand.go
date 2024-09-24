package cmd

import (
	"context"

	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize() //initialize stored creds, logging etc here later
	rootCmd.AddCommand(newVersionCmd().cmd)
	rootCmd.AddCommand(newBuildCmd().cmd)
	rootCmd.AddCommand(newTestCmd().cmd)
}

var rootCmd = &cobra.Command{
	Use:           "kinde",
	SilenceUsage:  true,
	SilenceErrors: true,
	Annotations:   map[string]string{},
	Version:       release.Branch,
	Short:         "Kinde CLI",
	Long:          "The official command-line for Kinde.",
}

func Execute(context context.Context) {
	if err := rootCmd.ExecuteContext(context); err != nil {
		rootCmd.PrintErr(err)
	}
}
