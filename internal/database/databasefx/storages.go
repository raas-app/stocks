package databasefx

import (
	raas "github.com/raas-app/stocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideDatabaseConnection(logger *zap.Logger, config *raas.Config) Connection {
	return NewConnectionBuilder(logger, config)
}

var Providers = fx.Module("databases",
	fx.Provide(ProvideDatabaseConnection),
)
