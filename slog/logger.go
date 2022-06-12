package slog

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Environment uint8

const (
	DevEnvironment Environment = iota
	ProdEnvironment
)

func Dev() (*zap.Logger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.EncoderConfig.TimeKey = ""

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to construct dev logger: %w", err)
	}

	return logger, nil
}

func Prod() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to construct prod logger: %w", err)
	}

	return logger, nil
}
