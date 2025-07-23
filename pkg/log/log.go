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
}

func GetSlogAdapter(l zerolog.Logger) *slog.Logger {
	logger := slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &l}.NewZerologHandler())
	return logger
}

func GetLogWriter(settings *SharedLogSettings) zerolog.Logger {

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	var zerologWriter zerolog.Logger

	prettyPrintLogs := true

	prettyPrintLogs, err := strconv.ParseBool(os.Getenv("KINDE_PRETTY_PRINT_LOGS"))
	if err != nil {
		prettyPrintLogs = settings.PrettyPrint
	}
	logLevel, ok := os.LookupEnv("KINDE_SERVER_LOG_LEVEL")
	if !ok {
		logLevel = "info"
	}

	zeroLogLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("Log Level of %v is not a valid value - Defaulting to INFO level\n", logLevel)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zeroLogLevel)
	}

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

		consoleWriter := zerolog.ConsoleWriter{
			Out:           os.Stderr,
			TimeFormat:    time.RFC3339,
			PartsOrder:    *settings.ConsolePartsOrder,
			PartsExclude:  *settings.ConsolePartsExclude,
			FieldsExclude: *settings.ConsoleFieldsExclude,
			//NoColor:       true,
			FormatExtra: settings.ConsoleFormatExtra,
		}

		//using local time for pretty logging
		consoleWriter.TimeFormat = time.Stamp
		consoleWriter.NoColor = false

		zerologWriter = zerolog.New(consoleWriter).With().Timestamp().Str("component", settings.ComponentName).Logger()
	} else {
		zerologWriter = zerolog.New(io.Writer(os.Stderr)).With().Timestamp().Str("component", settings.ComponentName).Logger()
	}

	return zerologWriter
}
