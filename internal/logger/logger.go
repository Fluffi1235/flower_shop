package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfgLogger := zap.NewProductionConfig()
	cfgLogger.DisableStacktrace = true

	logger, err := cfgLogger.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
