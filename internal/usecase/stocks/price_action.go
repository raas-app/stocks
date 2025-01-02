package stocksusecase

import (
	"context"

	raas "github.com/raas-app/stocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PriceActionUsecase interface {
	GetEodPriceAction(ctx context.Context, symbol string) []raas.EodPriceAction
	GetIntradayPriceAction(ctx context.Context, symbol string) []raas.IntradayPriceAction
}
type PriceActionUsecaseParams struct {
	fx.In

	StockFetcherHandler raas.StockFetcherHandler
	Logger              *zap.Logger
}

type priceActionUsecase struct {
	StockFetcherHandler raas.StockFetcherHandler
	Logger              *zap.Logger
}

func ProvidePriceActionUsecase(stockFetcherHandler raas.StockFetcherHandler, logger *zap.Logger) (PriceActionUsecase, error) {
	return &priceActionUsecase{
		StockFetcherHandler: stockFetcherHandler,
		Logger:              logger,
	}, nil
}

func (uc *priceActionUsecase) GetEodPriceAction(ctx context.Context, symbol string) []raas.EodPriceAction {
	return uc.StockFetcherHandler.GetEodPriceAction(ctx, symbol)
}

func (uc *priceActionUsecase) GetIntradayPriceAction(ctx context.Context, symbol string) []raas.IntradayPriceAction {
	return uc.StockFetcherHandler.GetIntradayPriceAction(ctx, symbol)
}
