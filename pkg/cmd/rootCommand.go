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

const authGroupID = "auth"

var errNoEnvironment = fmt.Errorf("no Kinde domain configured: pass --domain <your Kinde domain> or set %s", config.EnvDomain)

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

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	ctx = config.Ctx(ctx, cfg)

	rootCmd, err := NewRootCmd(ios, cfg)
	if err != nil {
		return err
	}
	rootCmd.SetArgs(args)

	updates := startUpdateCheck(ctx, ios, args)
	runErr := rootCmd.ExecuteContext(ctx)

	if err := cfg.SaveIfChanged(); err != nil {
		return errors.Join(runErr, fmt.Errorf("failed to save configuration: %w", err))
	}

	// Shown after the command's output, and not after an error, which should
	// be the last thing printed.
	if update := <-updates; update != nil && runErr == nil {
		fmt.Fprintf(ios.ErrOut, "\n%s\n", update.Message())
	}

	return runErr
}

// NewRootCmd builds the full command tree. Execute uses it to run the CLI;
// the docs generator uses it to write the command reference.
func NewRootCmd(ios *iostreams.IOStreams, cfg config.IConfig) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:           "kinde",
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations:   map[string]string{},
		// Any value enables --version; the template below prints the details.
		Version: release.String(),
		Short:   "Kinde CLI",
		Long:    "The official command-line for Kinde.",
	}
	rootCmd.SetVersionTemplate(release.String() + "\n")
	rootCmd.SetIn(ios.In)
	rootCmd.SetOut(ios.Out)
	rootCmd.SetErr(ios.ErrOut)

	var domainFlag string
	rootCmd.PersistentFlags().StringVar(&domainFlag, "domain", "",
		fmt.Sprintf("Kinde domain to use for this command (default: the domain you last logged in to; or set %s)", config.EnvDomain))

	resourceCmds, err := newResourceCommands(ios)
	if err != nil {
		return nil, err
	}
	accountCmd, err := newAccountCmd(ios)
	if err != nil {
		return nil, err
	}

	rootCmd.AddGroup(
		&cobra.Group{ID: authGroupID, Title: "Authentication:"},
		&cobra.Group{ID: resourceGroupID, Title: "Management API:"},
	)
	for _, authCmd := range []*cobra.Command{newLoginCmd(ios).cmd, newLogoutCmd(ios).cmd, accountCmd, newWhoAmI(ios).cmd} {
		authCmd.GroupID = authGroupID
		rootCmd.AddCommand(authCmd)
	}
	rootCmd.AddCommand(resourceCmds...)
	rootCmd.AddCommand(newVersionCmd(ios).cmd)

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

	return rootCmd, nil
}
