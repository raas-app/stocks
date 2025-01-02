package usecasefx

import (
	stocksusercase "github.com/raas-app/stocks/internal/usecase/stocks"
	"go.uber.org/fx"
)

var Providers = fx.Module("use-cases",
	fx.Provide(stocksusercase.ProvideStocksSymbolsUsecase),
	fx.Provide(stocksusercase.ProvidePriceActionUsecase),
)
