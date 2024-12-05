package zapper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapConfig represents zap.Logger configuration.
type ZapConfig struct {
	Config  zap.Config
	Options []zap.Option
}

// DefaultZapConfig return default configuration: format - json; level - warn; output - stdout;
// sampler: 100 unique messages per second, then every 100th message will be logged per same second.
func DefaultZapConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.WarnLevel),
		Encoding:    FormatJSON,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig:    DefaultZapEncoderConfig(),
		OutputPaths:      []string{OutputStdOut},
		ErrorOutputPaths: []string{OutputStdOut},
	}
}

// DefaultZapEncoderConfig returns encoder ZapConfig that suits company requirements.
func DefaultZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
