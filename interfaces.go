package raas

import (
	"context"
	"net/http"
)

type StockFetcherHandler interface {
	GetSymbols(ctx context.Context) []Stock
	GetIntradayPriceAction(ctx context.Context, symbol string) []IntradayPriceAction
	GetEodPriceAction(ctx context.Context, symbol string) []EodPriceAction
}

type SymbolsHttpHandler interface {
	GetSymbols(w http.ResponseWriter, r *http.Request)
}

type PriceActionHttpHandler interface {
	GetIntraday(w http.ResponseWriter, r *http.Request)
	GetEndOfDay(w http.ResponseWriter, r *http.Request)
}
