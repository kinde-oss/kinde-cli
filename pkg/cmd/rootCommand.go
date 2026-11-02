package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	cli_log "github.com/kinde-oss/kinde-cli/pkg/log"
	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/spf13/cobra"
)

var errNoEnvironment = errors.New("no Kinde domain configured. Run 'kinde login --domain <your Kinde domain>'")

// Execute runs the CLI with the given arguments (excluding the program name)
// and returns any error so the caller can set the exit code. All input and
// output goes through ios.
func Execute(ctx context.Context, args []string, ios *iostreams.IOStreams) error {

	logger := cli_log.GetLogWriter(&cli_log.SharedLogSettings{
		ComponentName:        "kinde_cli",
		PrettyPrint:          true,
		UseColor:             ios.ColorEnabled(),
		ConsoleFieldsExclude: &[]string{"component"}, //this is only used for pretty printing
		Out:                  ios.ErrOut,
	})
	log.Logger = logger
	zerolog.DefaultContextLogger = &log.Logger
	slog.SetDefault(cli_log.GetSlogAdapter(logger))
	ctx = logger.WithContext(ctx)

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
	rootCmd.SetIn(ios.In)
	rootCmd.SetOut(ios.Out)
	rootCmd.SetErr(ios.ErrOut)

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	ctx = config.Ctx(ctx, cfg)

	var domainFlag string
	rootCmd.PersistentFlags().StringVar(&domainFlag, "domain", "",
		fmt.Sprintf("Kinde domain to use for this command (default: the domain you last logged in to; or set %s)", config.EnvDomain))

	manageCmd, err := newManageCmd()
	if err != nil {
		return err
	}

	rootCmd.AddCommand(newVersionCmd(ios).cmd)
	rootCmd.AddCommand(newLoginCmd(ios).cmd)
	rootCmd.AddCommand(newLogoutCmd().cmd)
	rootCmd.AddCommand(newWhoAmI().cmd)
	rootCmd.AddCommand(manageCmd.cmd)

	// Runs after flags are parsed and before the command, so every command sees
	// the domain chosen for this run. Precedence: --domain, then KINDE_DOMAIN,
	// then the saved default.
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		domain := domainFlag
		if domain == "" {
			domain = os.Getenv(config.EnvDomain)
		}
		if domain != "" {
			if err := cfg.UseDomain(domain); err != nil {
				return err
			}
		}

		if err, isValid := cfg.Validate(); !isValid {
			log.Error().Err(err).Msg("Issues with configuration")
		}
		return nil
	}

	runErr := rootCmd.ExecuteContext(ctx)

	if err := cfg.SaveIfChanged(); err != nil {
		return errors.Join(runErr, fmt.Errorf("failed to save configuration: %w", err))
	}

	return runErr
}
