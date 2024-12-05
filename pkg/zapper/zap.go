// Package logger provides handy methods to init zap logger according company requirements.
package zapper

import (
	"log"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZap creates instance of zap.Logger.
func NewZap(options ...Option) (*zap.Logger, error) {
	cfg, err := newZap(options...)
	if err != nil {
		return nil, err
	}
	return cfg.Config.Build(cfg.Options...)
}

// NewZapWithLevel creates instance of zap.Logger and returns *zap.AtomicLevel to check and change log level.
func NewZapWithLevel(options ...Option) (*zap.Logger, *zap.AtomicLevel, error) {
	cfg, err := newZap(options...)
	if err != nil {
		return nil, nil, err
	}
	l, err := cfg.Config.Build(cfg.Options...)
	if err != nil {
		return nil, nil, err
	}
	return l, &cfg.Config.Level, nil
}

// NewZapWithHandler creates instance of zap.Logger and returns http.Handler to check and change log level.
func NewZapWithHandler(options ...Option) (*zap.Logger, http.Handler, error) {
	return NewZapWithLevel(options...)
}

// newZap build zap logger configured with options.
func newZap(options ...Option) (*ZapConfig, error) {
	cfg := &ZapConfig{
		Config:  DefaultZapConfig(),
		Options: make([]zap.Option, 0),
	}
	for _, opt := range options {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	if !cfg.Config.DisableCaller {
		cfg.Options = append(cfg.Options, zap.AddCallerSkip(1), zap.AddCaller())
	}
	return cfg, nil
}

// NewSugaredZap creates instance of zap.SugaredLogger. Please use this method
// only if your app uses zap.SugaredLogger by default, in other case please use
// zap.Logger.Sugar() method explicitly.
func NewSugaredZap(options ...Option) (*zap.SugaredLogger, error) {
	l, err := NewZap(options...)
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}

// NewStdLogger creates log.Logger instance from zap.Logger.
func NewStdLogger(l *zap.Logger) (*log.Logger, error) {
	return zap.NewStdLogAt(l, l.Level())
}

// DebugEnabled returns true if logger level is debug.
func DebugEnabled(log ZapLevelLogger) bool {
	return log.Level() == zapcore.DebugLevel
}

// ZapLevelLogger interface required to check level both for zap.Logger and zap.SugaredLogger.
type ZapLevelLogger interface {
	// Level return current log level.
	Level() zapcore.Level
}

// SetLevel change log level.
func SetLevel(l *zap.AtomicLevel, level string) error {
	return l.UnmarshalText([]byte(level))
}
