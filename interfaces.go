package raas

import (
	"context"
	"net/http"
)

type StockFetcherHandler interface {
	GetSymbols(ctx context.Context) []Stock
	GetIntradayPriceAction(symbol string) []IntradayPriceAction
	GetEodPriceAction(symbol string) []EodPriceAction
}

type SymbolsHttpHandler interface {
	GetSymbols(w http.ResponseWriter, r *http.Request)
}
