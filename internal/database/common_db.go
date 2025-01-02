package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// commonDB обертка Common, сделан для отправки метрики в prometheus
type commonDB struct {
	Common
	// alias Псевдоним базы данных
	alias  string
	logger *zap.Logger
}

func NewCommonDB(db *sqlx.DB, alias string, logger *zap.Logger) (*commonDB, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}
	cDb := commonDB{
		Common: db,
		alias:  alias,
		logger: logger,
	}
	return &cDb, nil
}

func (c commonDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := c.Common.QueryContext(ctx, query, args...)
	return rows, err
}

func (c commonDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	row := c.Common.QueryRowContext(ctx, query, args...)
	return row
}

func (c commonDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	res, err := c.Common.ExecContext(ctx, query, args...)
	return res, err
}
