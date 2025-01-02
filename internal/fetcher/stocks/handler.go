package stocks

import (
	raas "github.com/raas-app/stocks"
	"go.uber.org/zap"
)

type PriceActionResponse struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    [][]float64 `json:"data"`
}

type stockHandler struct {
	Config *raas.Config
	Logger *zap.Logger
}

func NewStockHandler(config *raas.Config, logger *zap.Logger) raas.StockFetcherHandler {
	return &stockHandler{
		Config: config,
		Logger: logger,
	}
}

func InitializeStockHandler(h raas.StockFetcherHandler) error {
	h.GetEodPriceAction("EFERT")
	return nil
}
