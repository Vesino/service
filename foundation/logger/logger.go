// Package logger provides a convenience function to constructing a logger
// for use. This is required not just for applications but for testing.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New constructs a Sugared Logger that writes to stdout and
// provides human-readable timestamps.
func New(service string, outputPaths ...string) (*zap.SugaredLogger, error) {
	zap.NewProductionConfig()
	zapcore.New()
}
