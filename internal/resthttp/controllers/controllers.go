package controllers

import (
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	priceaction "github.com/raas-app/stocks/internal/resthttp/controllers/price-action"
	"github.com/raas-app/stocks/internal/resthttp/controllers/symbols"
	stocksusecase "github.com/raas-app/stocks/internal/usecase/stocks"
	"go.uber.org/zap"
)

type Controller struct {
	SymbolHandler      func(*respond.Responder, *zap.Logger, stocksusecase.SymbolsUsecase) (raas.SymbolsHttpHandler, error)
	PriceActionHandler func(*respond.Responder, *zap.Logger, stocksusecase.PriceActionUsecase) (raas.PriceActionHttpHandler, error)
}

func Controllers() *Controller {
	return &Controller{
		SymbolHandler:      symbols.NewSymbolsHandler,
		PriceActionHandler: priceaction.NewPriceActionHandler,
	}
}
