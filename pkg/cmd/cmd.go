package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "kinde",
	SilenceUsage:  true,
	SilenceErrors: true,
	Annotations:   map[string]string{},
	Version:       "master",
	Short:         "CLI to assist you integrate and work with Kinde",
	Long:          "The official command-line for Kinde.",
}

func Execute(context context.Context) {
	if err := rootCmd.ExecuteContext(context); err != nil {
	}
}

func init() {
	cobra.OnInitialize() //initialize stored creds, logging etc here later

}
