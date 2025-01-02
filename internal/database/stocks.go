package database

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raas-app/stocks/internal/database/models"
	"go.uber.org/zap"
)

type StockStore interface {
	GetStocks(ctx context.Context) ([]models.Stock, error)
}

type stocksStorage struct {
	db     Common
	logger *zap.Logger
}

func NewStocksStorage(db Common, logger *zap.Logger) *stocksStorage {

	return &stocksStorage{
		db:     db,
		logger: logger,
	}
}

func (s *stocksStorage) GetStocks(ctx context.Context) ([]models.Stock, error) {
	query := "SELECT `symbol` FROM `stock`"

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		s.logger.Error("failed to get stocks", zap.Error(err))
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			s.logger.Error("failed to close stocks rows", zap.Error(err))
		}
	}(rows)
	stocks := make([]models.Stock, 0)
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(
			&stock.Symbol,
		)
		if err != nil {
			s.logger.Error("failed to scan stock", zap.Error(err))
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
