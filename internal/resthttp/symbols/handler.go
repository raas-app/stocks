package symbols

import (
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	"go.uber.org/zap"
)

type SymbolsHandler struct {
	responder           *respond.Responder
	logger              *zap.Logger
	stockFetcherHandler raas.StockFetcherHandler
}

func NewSymbolsHandler(responder *respond.Responder, logger *zap.Logger, stockHandler raas.StockFetcherHandler) (raas.SymbolsHttpHandler, error) {

	return &SymbolsHandler{
		responder:           responder,
		logger:              logger,
		stockFetcherHandler: stockHandler,
	}, nil
}
