package symbols

import (
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	stocksusecase "github.com/raas-app/stocks/internal/usecase/stocks"
	"go.uber.org/zap"
)

type SymbolsHandler struct {
	responder      *respond.Responder
	logger         *zap.Logger
	SymbolsUsecase stocksusecase.SymbolsUsecase
}

func NewSymbolsHandler(responder *respond.Responder, logger *zap.Logger, symbolsUsecase stocksusecase.SymbolsUsecase) (raas.SymbolsHttpHandler, error) {

	return &SymbolsHandler{
		responder:      responder,
		logger:         logger,
		SymbolsUsecase: symbolsUsecase,
	}, nil
}
