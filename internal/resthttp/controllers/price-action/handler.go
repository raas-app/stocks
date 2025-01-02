package priceaction

import (
	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	stocksusecase "github.com/raas-app/stocks/internal/usecase/stocks"
	"go.uber.org/zap"
)

type PriceActionHandler struct {
	responder *respond.Responder
	logger    *zap.Logger
	usecase   stocksusecase.PriceActionUsecase
}

func NewPriceActionHandler(responder *respond.Responder, logger *zap.Logger, usecase stocksusecase.PriceActionUsecase) (raas.PriceActionHttpHandler, error) {
	return &PriceActionHandler{
		responder: responder,
		logger:    logger,
		usecase:   usecase,
	}, nil
}
