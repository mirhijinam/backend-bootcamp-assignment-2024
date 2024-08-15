package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(mode, filepath string) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := zap.NewProductionConfig()
	var lvl zapcore.Level
	var encoding string
	switch mode {
	case "info":
		lvl = zap.InfoLevel
		encoding = "console"
	case "debug":
		lvl = zap.DebugLevel
		encoding = "json"
	}

	cfg.Level = zap.NewAtomicLevelAt(lvl)
	cfg.Encoding = encoding
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{
		filepath, "stderr",
	}
	cfg.ErrorOutputPaths = []string{
		filepath, "stderr",
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to build logger: %v", err))
	}

	return logger
}
