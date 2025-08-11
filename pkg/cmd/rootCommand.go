package cmd

import (
	"context"
	"log/slog"

	"github.com/jwalton/go-supportscolor"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	cli_log "github.com/kinde-oss/kinde-cli/pkg/log"
	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()

	log.Logger = cli_log.GetLogWriter(&cli_log.SharedLogSettings{
		ComponentName:        "kinde_cli",
		PrettyPrint:          true,
		UseColor:             supportscolor.Stdout().SupportsColor,
		ConsoleFieldsExclude: &[]string{"component"}, //this is only used for pretty printing
	})

	zerolog.DefaultContextLogger = &log.Logger
	slog.SetDefault(cli_log.GetSlogAdapter(log.Logger))
}

func Execute(ctx context.Context) {

	rootCmd := &cobra.Command{
		Use:           "kinde",
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations:   map[string]string{},
		Version:       release.Branch,
		Short:         "Kinde CLI",
		Long:          "The official command-line for Kinde.",
	}

	cfg, err := config.NewConfig()
	ctx = config.Ctx(ctx, cfg)

	rootCmd.PersistentFlags().Func("domain", "Specific kinde domain to authenticate against.", func(val string) error {
		return cfg.SwitchEnvironment(val)
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize configuration")
		return
	}

	rootCmd.AddCommand(newVersionCmd().cmd)
	//rootCmd.AddCommand(newBuildCmd().cmd)
	rootCmd.AddCommand(newLoginCmd().cmd)
	rootCmd.AddCommand(newLogoutCmd().cmd)
	rootCmd.AddCommand(newWhoAmI().cmd)
	rootCmd.AddCommand(newManageCmd(ctx).cmd)

	cfg.PersistConfig()

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err, isValid := cfg.Validate(); !isValid {
			log.Error().Err(err).Msg("Issues with configuration")
		}
		return nil
	}

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Error().Err(err).Msg("Command execution failed")
	}

	if err := cfg.PersistConfig(); err != nil {
		log.Error().Err(err).Msg("Failed to persist configuration")
	}
}
