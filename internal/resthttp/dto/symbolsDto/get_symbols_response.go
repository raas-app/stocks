package symbolsdto

import (
	"github.com/raas-app/stocks/internal/database/models"
)

type GetSymbolsResponse struct {
	Symbols []string `json:"symbols"`
	Count   int      `json:"count"`
}

func NewGetSymbolsResponse(stocks []models.Stock) *GetSymbolsResponse {
	var symbols = make([]string, 0)
	for _, stock := range stocks {
		symbols = append(symbols, stock.Symbol)
	}
	return &GetSymbolsResponse{
		Symbols: symbols,
		Count:   len(symbols),
	}
}
