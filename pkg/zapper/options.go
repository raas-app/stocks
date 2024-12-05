package zapper

import (
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug   = "debug"
	LevelInfo    = "info"
	LevelWarn    = "warn"
	levelWarning = "warning"
	LevelError   = "error"

	FormatJSON    = "json"
	FormatConsole = "console"
	FormatColor   = "color"

	OutputStdOut = "stdout"
	OutputStdErr = "stderr"
)

// Option allows to configure logger.
type Option func(p *ZapConfig) error

// WithLevel set the minimum enabled logging level.
// Valid level values (case insensitive): debug|info|warn|error|dpanic|panic|fatal
func WithLevel(level string) Option {
	level = strings.ToLower(level)
	// some other services have `warn` level as `WARNING`
	if level == levelWarning {
		level = LevelWarn
	}

	return func(cfg *ZapConfig) error {
		if level == "" {
			return errors.New("log level is empty")
		}
		if err := cfg.Config.Level.UnmarshalText([]byte(level)); err != nil {
			return fmt.Errorf("invalid logger level `%s`: %w", level, err)
		}
		return nil
	}
}

// WithAtomicLevel sets the atomic logging level.
func WithAtomicLevel(level zap.AtomicLevel) Option {
	return func(cfg *ZapConfig) error {
		cfg.Config.Level = level
		return nil
	}
}

// WithOutputFormat set output format.
// Valid options (case-insensitive): json|text|console|color|colour
func WithOutputFormat(format string) Option {
	return func(cfg *ZapConfig) error {
		switch strings.ToLower(format) {
		case FormatJSON:
			cfg.Config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
			cfg.Config.Encoding = FormatJSON
		case FormatConsole, "text":
			cfg.Config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
			cfg.Config.Encoding = FormatConsole
		case FormatColor, "colour":
			cfg.Config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			cfg.Config.Encoding = FormatConsole
		default:
			return fmt.Errorf("unknown log format: `%s`", format)
		}
		return nil
	}
}

// WithDevelopment enables panic for DPanic log level and stack traces for warn level.
func WithDevelopment() Option {
	return func(cfg *ZapConfig) error {
		cfg.Config.Development = true
		return nil
	}
}

// WithDisabledCaller disables information about caller
func WithDisabledCaller() Option {
	return func(cfg *ZapConfig) error {
		cfg.Config.DisableCaller = true
		return nil
	}
}

// WithDisabledStackTrace disables stacktraces
func WithDisabledStackTrace() Option {
	return func(cfg *ZapConfig) error {
		cfg.Config.DisableStacktrace = true
		return nil
	}
}

// WithOutputPaths is a list of URLs or file paths to write logging output to.
func WithOutputPaths(paths ...string) Option {
	return func(cfg *ZapConfig) error {
		if len(paths) > 0 {
			cfg.Config.OutputPaths = paths
			cfg.Config.ErrorOutputPaths = paths
		}
		return nil
	}
}

// WithOptions passes zap.Option to constructor.
func WithOptions(options ...zap.Option) Option {
	return func(cfg *ZapConfig) error {
		if len(options) > 0 {
			cfg.Options = append(cfg.Options, options...)
		}
		return nil
	}
}
