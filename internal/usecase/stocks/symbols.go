package stocksusecase

import (
	"context"
	"fmt"

	raas "github.com/raas-app/stocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SymbolsUsecase interface {
	GetSymbols(ctx context.Context) []raas.Stock
}

type SymbolsUsecaseParams struct {
	fx.In

	StockFetcherHandler raas.StockFetcherHandler
	Logger              *zap.Logger
}

type symbolsUsecase struct {
	stockFetcherHandler raas.StockFetcherHandler
	logger              *zap.Logger
}

func ProvideStocksSymbolsUsecase(p SymbolsUsecaseParams) (SymbolsUsecase, error) {
	if p.Logger == nil {
		return nil, fmt.Errorf("missing Logger dependencies")
	}
	if p.StockFetcherHandler == nil {
		return nil, fmt.Errorf("missing StockFetcherHandler dependencies")
	}

	return &symbolsUsecase{
		stockFetcherHandler: p.StockFetcherHandler,
		logger:              p.Logger,
	}, nil
}

func (uc *symbolsUsecase) GetSymbols(ctx context.Context) []raas.Stock {
	return uc.stockFetcherHandler.GetSymbols(ctx)
}
