package log

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/UnnoTed/horizontal"
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
	logLevel := zerolog.InfoLevel

	if structuredLog, err := strconv.ParseBool(os.Getenv("KINDE_STRUCTURED_LOG")); err == nil {
		prettyPrintLogs = !structuredLog
	}

	if logLevel, ok := os.LookupEnv("KINDE_LOG_LEVEL"); ok {
		logLevel, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			fmt.Printf("Log Level of %v is not a valid value - Defaulting to INFO level\n", logLevel)
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

	}

	zerolog.SetGlobalLevel(logLevel)

	if prettyPrintLogs {

		consoleWriter := horizontal.ConsoleWriter{Out: os.Stderr}

		zerologWriter = zerolog.New(consoleWriter).With().Timestamp().Str("component", settings.ComponentName).Logger()
	} else {
		zerologWriter = zerolog.New(io.Writer(os.Stderr)).With().Timestamp().Str("component", settings.ComponentName).Logger()
	}

	return zerologWriter
}
