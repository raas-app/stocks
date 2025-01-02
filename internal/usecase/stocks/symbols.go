package stocksusecase

import (
	"context"
	"fmt"

	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/database"
	"github.com/raas-app/stocks/internal/database/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SymbolsUsecase interface {
	GetSymbols(ctx context.Context) ([]models.Stock, error)
}

type SymbolsUsecaseParams struct {
	fx.In

	StockFetcherHandler raas.StockFetcherHandler
	StockStore          database.StockStore
	Logger              *zap.Logger
}

type symbolsUsecase struct {
	stockFetcherHandler raas.StockFetcherHandler
	StockStore          database.StockStore
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
		StockStore:          p.StockStore,
		logger:              p.Logger,
	}, nil
}

func (uc *symbolsUsecase) GetSymbols(ctx context.Context) ([]models.Stock, error) {
	stocks, err := uc.StockStore.GetStocks(ctx)
	if err != nil {
		uc.logger.Error("failed to get stocks", zap.Error(err))
		return nil, err
	}
	return stocks, nil
}
