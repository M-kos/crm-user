package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	Log *slog.Logger
}

func NewLogger() *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return &Logger{
		Log: logger,
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.Log.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Log.Error(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.Log.Warn(msg, args...)
}
