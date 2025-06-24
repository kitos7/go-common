package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type contextKey string

const loggerKey = contextKey("logger")

// WithLogger returns a new context with the logger
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext retrieves a logger from the context or returns the default logger if not found
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

// InfoKV logs an informational message with key-value pairs
func InfoKV(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Info(msg, args...)
}

// ErrorKV logs an error message with key-value pairs
func ErrorKV(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(msg, args...)
}

// DebugKV logs a debug message with key-value pairs
func DebugKV(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Debug(msg, args...)
}

// WarnKV logs a warning message with key-value pairs
func WarnKV(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Warn(msg, args...)
}

// FatalKV logs an error message with key-value pairs and then exits the program
func FatalKV(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(msg, args...)
	os.Exit(1)
}

// Infof logs an informational message with printf-style formatting
func Infof(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Info(fmt.Sprintf(msg, args...))
}

// Errorf logs an error message with printf-style formatting
func Errorf(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(fmt.Sprintf(msg, args...))
}

// Debugf logs a debug message with printf-style formatting
func Debugf(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Debug(fmt.Sprintf(msg, args...))
}

// Warnf logs a warning message with printf-style formatting
func Warnf(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Warn(fmt.Sprintf(msg, args...))
}

// Fatalf logs an error message with printf-style formatting and then exits the program
func Fatalf(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(fmt.Sprintf(msg, args...))
	os.Exit(1)
}
