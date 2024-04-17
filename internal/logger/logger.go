package logger

import (
	"log/slog"
	"os"
)

const (
	devEnv  = "dev"
	prodEnv = "prod"
)

func SetupNew(environment string) *slog.Logger {
	var log *slog.Logger

	switch environment {
	case devEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prodEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
