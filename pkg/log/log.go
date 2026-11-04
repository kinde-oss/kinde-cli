package log

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
)

type SharedLogSettings struct {
	ComponentName        string
	ConsolePartsExclude  *[]string
	ConsoleFieldsExclude *[]string
	ConsolePartsOrder    *[]string
	ConsoleFormatExtra   func(m map[string]interface{}, buf *bytes.Buffer) error
	PrettyPrint          bool
	UseColor             bool
	// Out is where log lines are written; defaults to os.Stderr.
	Out io.Writer
}

func GetSlogAdapter(l zerolog.Logger) *slog.Logger {
	logger := slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &l}.NewZerologHandler())
	return logger
}

func GetLogWriter(settings *SharedLogSettings) zerolog.Logger {

	out := settings.Out
	if out == nil {
		out = os.Stderr
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(out, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})))

	prettyPrintLogs := true
	// Logs are for diagnosing problems; commands print their own output.
	// Set KINDE_LOG_LEVEL=debug or info to see more.
	logLevel := zerolog.WarnLevel

	if structuredLog, err := strconv.ParseBool(os.Getenv("KINDE_STRUCTURED_LOG")); err == nil {
		prettyPrintLogs = !structuredLog
	}

	if logLevelStr := os.Getenv("KINDE_LOG_LEVEL"); logLevelStr != "" {
		parsedLogLevel, err := zerolog.ParseLevel(logLevelStr)
		if err != nil {
			fmt.Fprintf(out, "KINDE_LOG_LEVEL %q is not a valid level (debug, info, warn, error); using warn\n", logLevelStr)
		} else {
			logLevel = parsedLogLevel
		}
	}

	zerolog.SetGlobalLevel(logLevel)

	var zerologWriter zerolog.Logger

	if prettyPrintLogs {

		if settings.ConsolePartsExclude == nil {
			settings.ConsolePartsExclude = &[]string{"sub_component", "elapsed"}
		}

		if settings.ConsoleFieldsExclude == nil {
			settings.ConsoleFieldsExclude = &[]string{}
		}

		if settings.ConsolePartsOrder == nil {
			settings.ConsolePartsOrder = &[]string{
				zerolog.TimestampFieldName,
				"component",
				zerolog.LevelFieldName,
				zerolog.CallerFieldName,
				zerolog.MessageFieldName,
			}
		}

		consoleWriter := ConsoleWriter{
			Out:           out,
			TimeFormat:    time.RFC3339,
			PartsOrder:    *settings.ConsolePartsOrder,
			PartsExclude:  *settings.ConsolePartsExclude,
			FieldsExclude: *settings.ConsoleFieldsExclude,
			NoColor:       !settings.UseColor,
			FormatExtra: func(m map[string]any, buf *bytes.Buffer) error {
				return nil
			},
		}

		//using local time for pretty logging
		consoleWriter.TimeFormat = time.Stamp

		zerologWriter = zerolog.New(consoleWriter).With().Timestamp().Str("component", settings.ComponentName).Logger()
	} else {
		zerologWriter = zerolog.New(out).With().Timestamp().Str("component", settings.ComponentName).Logger()
	}

	return zerologWriter
}
