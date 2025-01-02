package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/database/databasefx"
	"github.com/raas-app/stocks/internal/fetcher/fetcherfx"
	"github.com/raas-app/stocks/internal/resthttp"
	"github.com/raas-app/stocks/internal/scrapper/scrapperfx"
	"github.com/raas-app/stocks/internal/usecase/usecasefx"
	"github.com/raas-app/stocks/pkg/zapper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultStartTimeout = 15 * time.Second
	defaultStopTimeout  = 15 * time.Second
)

var (
	serviceStartTimeout, serviceStopTimeout time.Duration
)

func getStartTimeout() time.Duration {
	if serviceStartTimeout > 1*time.Second {
		return serviceStartTimeout
	}
	return defaultStartTimeout
}

func getStopTimeout() time.Duration {
	if serviceStopTimeout > 1*time.Second {
		return serviceStopTimeout
	}
	return defaultStopTimeout
}

func envName() raas.EnvName {
	name := os.Getenv("ENV_NAME")
	if name == "" {
		name = "undefined"
	}
	return raas.EnvName(name)
}

func appName() raas.AppName {
	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "raas-stocks"
	}
	return raas.AppName(name)
}

func newLogger(level string, debug, console bool) (*zap.Logger, error) {
	options := []zapper.Option{
		zapper.WithLevel(level),
		zapper.WithDisabledCaller(),
		func(zc *zapper.ZapConfig) error {
			zc.Config.EncoderConfig.TimeKey = "@timestamp"
			zc.Config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
			return nil
		},
	}

	if debug {
		options = append(options, zapper.WithDevelopment())
		if console {
			options = append(options, zapper.WithOutputFormat(zapper.FormatConsole))
		}
		options = append(options, func(zc *zapper.ZapConfig) error {
			zc.Config.Sampling = nil
			return nil
		})
	} else {
		samplingConfig := zap.SamplingConfig{
			Initial:    250,
			Thereafter: 250,
		}
		options = append(options, func(zc *zapper.ZapConfig) error {
			zc.Config.Sampling = &zap.SamplingConfig{
				Initial:    samplingConfig.Initial,
				Thereafter: samplingConfig.Thereafter,
			}
			return nil
		})
	}

	logger, err := zapper.NewZap(options...)
	return logger, err
}

func appLogger(cfg *raas.Config, lc fx.Lifecycle) (*zap.Logger, error) {
	console := os.Getenv("ORDER_LOGGER") == "console"
	logger, err := newLogger("INFO", true, console)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}
	logger.Info("Config loaded successfully", zap.Any("config", cfg))

	lc.Append(fx.StopHook(logger.Sync))

	return logger, nil
}

var initApp = fx.Module("raas-app",
	fx.Provide(raas.Load, appLogger, envName, appName),
	fx.Invoke(func(cfg *raas.Config, log *zap.Logger) {
		log.Info("Application is starting", zap.Int("port", cfg.Server.Port))
		log.Info("Application configuration", zap.Any("config", cfg))
	}))

func main() {
	fx.New(
		initApp,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		scrapperfx.Providers,
		fetcherfx.Providers,
		databasefx.Providers,
		usecasefx.Providers,
		resthttp.Providers,
		databasefx.Module,
		resthttp.Launcher,
		fx.StartTimeout(getStartTimeout()),
		fx.StopTimeout(getStopTimeout()),
	).Run()
}
