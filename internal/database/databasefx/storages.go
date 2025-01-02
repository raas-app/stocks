package databasefx

import (
	"github.com/raas-app/stocks/internal/database"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideStocksStorage(logger *zap.Logger, db database.Database) database.StockStore {
	return database.NewStocksStorage(db.RW, logger)
}

var Providers = fx.Module("storages", fx.Provide(
	ProvideStocksStorage,
))
