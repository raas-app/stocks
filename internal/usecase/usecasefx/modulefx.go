package usecasefx

import (
	"github.com/raas-app/stocks/internal/usecase/stocks"
	"go.uber.org/fx"
)

var Providers = fx.Module("use-cases",
	fx.Provide(stocks.ProvideStocksSymbolsUsecase))
