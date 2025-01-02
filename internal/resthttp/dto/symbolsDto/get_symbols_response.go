package symbolsdto

import raas "github.com/raas-app/stocks"

type GetSymbolsResponse struct {
	Symbols []string `json:"symbols"`
	Count   int      `json:"count"`
}

func NewGetSymbolsResponse(stocks []raas.Stock) *GetSymbolsResponse {
	var symbols = make([]string, 0)
	for _, stock := range stocks {
		symbols = append(symbols, stock.Symbol)
	}
	return &GetSymbolsResponse{
		Symbols: symbols,
		Count:   len(symbols),
	}
}
