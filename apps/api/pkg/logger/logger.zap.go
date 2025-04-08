package logger

import (
	"github.com/esc-chula/intania-vote/apps/api/pkg/config"
	"go.uber.org/zap"
)

const (
	DEV  = "development"
	PROD = "production"
)

func NewLogger(cfg *config.Config) *zap.Logger {
	logger := newLoggerFactory(cfg.Environment)
	zap.ReplaceGlobals(logger)
	return logger
}

func newLoggerFactory(env string) *zap.Logger {
	switch env {
	case DEV:
		return zap.Must(zap.NewDevelopment())
	case PROD:
		return zap.Must(zap.NewProduction())
	default:
		return nil
	}
}
