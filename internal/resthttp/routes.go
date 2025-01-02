package resthttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	"github.com/raas-app/stocks/internal/resthttp/middlewares"
	"github.com/raas-app/stocks/internal/resthttp/symbols"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RouterDependencies struct {
	fx.In

	Logger              *zap.Logger
	StockFetcherHandler raas.StockFetcherHandler
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

	symbolsHandler, err := symbols.NewSymbolsHandler(
		responder,
		routerDependencies.Logger,
		routerDependencies.StockFetcherHandler,
	)
	if err != nil {
		return nil, err
	}
	router.Route("/api/v1/stocks", func(r chi.Router) {
		router.Use(debugLogger.LogRequest)
		r.Route("/symbols", func(r chi.Router) {
			r.Get("/", symbolsHandler.GetSymbols)
		})
	})
	return router, nil
}
