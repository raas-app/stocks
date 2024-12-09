package fetcherfx

import (
	raas "github.com/raas-app/stocks"
	fetcher "github.com/raas-app/stocks/internal/fetcher/stocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideStockFetcher(config *raas.Config, logger *zap.Logger) (fetcher.StockHandler, error) {
	return fetcher.NewStockHandler(config, logger), nil
}

var Providers = fx.Module("fetchers",
	fx.Provide(ProvideStockFetcher))
