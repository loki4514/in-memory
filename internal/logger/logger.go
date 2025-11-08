package logger

import (
	"os"

	"github.com/loki4514/in-memory.git/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewLogger(cfg *config.Config) zerolog.Logger {
	var logLevel zerolog.Level
	switch cfg.Logging.Level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if cfg.Logging.Format == "console" {
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat}).
			Level(logLevel).With().Timestamp().Logger()
	}

	return zerolog.New(os.Stdout).
		Level(logLevel).With().Timestamp().Logger()
}
