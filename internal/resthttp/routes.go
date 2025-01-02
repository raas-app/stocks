package resthttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	"github.com/raas-app/stocks/internal/resthttp/controllers"
	"github.com/raas-app/stocks/internal/resthttp/middlewares"
	stocksusecase "github.com/raas-app/stocks/internal/usecase/stocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RouterDependencies struct {
	fx.In

	Logger              *zap.Logger
	StockFetcherHandler raas.StockFetcherHandler
	SymbolsUsecase      stocksusecase.SymbolsUsecase
	PriceActionUsecase  stocksusecase.PriceActionUsecase
}

func MakeRoutes(routerDependencies RouterDependencies) (http.Handler, error) {
	responder, err := respond.NewResponder()
	if err != nil {
		return nil, err
	}

	debugLogger, err := middlewares.NewDebugLogger(routerDependencies.Logger)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	controllers := controllers.Controllers()

	symbolsHandler, err := controllers.SymbolHandler(
		responder,
		routerDependencies.Logger,
		routerDependencies.SymbolsUsecase,
	)
	if err != nil {
		return nil, err
	}

	priceActionHandler, err := controllers.PriceActionHandler(
		responder,
		routerDependencies.Logger,
		routerDependencies.PriceActionUsecase,
	)
	if err != nil {
		return nil, err
	}
	router.Route("/api/v1/stocks", func(r chi.Router) {
		router.Use(debugLogger.LogRequest)
		r.Route("/symbols", func(r chi.Router) {
			r.Get("/", symbolsHandler.GetSymbols)
		})
		r.Route("/price-action", func(r chi.Router) {
			r.Get("/{symbol}/intraday", priceActionHandler.GetIntraday)
			r.Get("/{symbol}/eod", priceActionHandler.GetEndOfDay)
		})
	})
	return router, nil
}
