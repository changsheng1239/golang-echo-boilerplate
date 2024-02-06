package log

import (
	"log/slog"
	"os"

	"github.com/charmbracelet/log"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type Logger struct {
	*slog.Logger
}

func NewLog(appEnv string) *Logger {
	return initLogger(appEnv)
}

func initLogger(appEnv string) *Logger {
	logger := log.New(os.Stdout)
	switch appEnv {
	case EnvDevelopment:
		logger.SetReportCaller(true)
		logger.SetLevel(log.DebugLevel)
	case EnvProduction:
		logger.SetLevel(log.InfoLevel)
		logger.SetFormatter(log.JSONFormatter)
	default:
		logger.SetLevel(log.InfoLevel)
	}

	slogger := slog.New(logger)
	return &Logger{slogger}
}
