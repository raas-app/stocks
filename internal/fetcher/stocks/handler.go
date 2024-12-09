package stocks

import (
	raas "github.com/raas-app/stocks"
	"go.uber.org/zap"
)

type StockHandler interface {
	GetSymbols() []Stock
	GetIntradayPriceAction(symbol string) []IntradayPriceAction
	GetEodPriceAction(symbol string) []EodPriceAction
}

type PriceActionResponse struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    [][]float64 `json:"data"`
}
type IntradayPriceAction struct {
	Price  float64
	Volume float64
	Time   float64
}

type EodPriceAction struct {
	Time   float64
	Open   float64
	Close  float64
	Volume float64
}
type Stock struct {
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	SectorName string `json:"sectorName"`
	IsETF      bool   `json:"isETF"`
	IsDebt     bool   `json:"isDebt"`
}

type stockHandler struct {
	Config *raas.Config
	Logger *zap.Logger
}

func NewStockHandler(config *raas.Config, logger *zap.Logger) StockHandler {
	return &stockHandler{
		Config: config,
		Logger: logger,
	}
}

func InitializeStockHandler(h StockHandler) error {
	h.GetEodPriceAction("EFERT")
	return nil
}
