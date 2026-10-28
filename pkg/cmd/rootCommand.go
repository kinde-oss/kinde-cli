package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jwalton/go-supportscolor"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	cli_log "github.com/kinde-oss/kinde-cli/pkg/log"
	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/spf13/cobra"
)

var errNoEnvironment = errors.New("no Kinde domain configured. Run 'kinde login --domain <your Kinde domain>'")

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

// Execute runs the CLI with the given arguments (excluding the program name)
// and returns any error so the caller can set the exit code.
func Execute(ctx context.Context, args []string) error {

	rootCmd := &cobra.Command{
		Use:           "kinde",
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations:   map[string]string{},
		Version:       release.Version,
		Short:         "Kinde CLI",
		Long:          "The official command-line for Kinde.",
	}
	rootCmd.SetArgs(args)

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	ctx = config.Ctx(ctx, cfg)

	rootCmd.PersistentFlags().Func("domain", "Kinde domain (environment) to work against.", func(val string) error {
		return cfg.SwitchEnvironment(val)
	})

	manageCmd, err := newManageCmd(ctx)
	if err != nil {
		return err
	}

	rootCmd.AddCommand(newVersionCmd().cmd)
	rootCmd.AddCommand(newLoginCmd().cmd)
	rootCmd.AddCommand(newLogoutCmd().cmd)
	rootCmd.AddCommand(newWhoAmI().cmd)
	rootCmd.AddCommand(manageCmd.cmd)

	cfg.PersistConfig()

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err, isValid := cfg.Validate(); !isValid {
			log.Error().Err(err).Msg("Issues with configuration")
		}
		return nil
	}

	runErr := rootCmd.ExecuteContext(ctx)

	if err := cfg.PersistConfig(); err != nil {
		return errors.Join(runErr, fmt.Errorf("failed to save configuration: %w", err))
	}

	return runErr
}
