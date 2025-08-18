package infra

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func initLogger() {
	once.Do(func() {
		logFile, err := os.OpenFile("app.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			slog.Default().Error("could not open log file", "error", err)
			logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
			return
		}
		logger = slog.New(slog.NewJSONHandler(logFile, nil))
	})
}

// Info logs an informational message
func Info(msg string, args ...any) {
	initLogger()
	logger.Info(msg, args...)
}

// Error logs an error message
func Error(msg string, args ...any) {
	initLogger()
	logger.Error(msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...any) {
	initLogger()
	logger.Warn(msg, args...)
}

// Debug logs a debug message
func Debug(msg string, args ...any) {
	initLogger()
	logger.Debug(msg, args...)
}
