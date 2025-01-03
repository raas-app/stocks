package models

type StockIndexMapping struct {
	ID          int    `db:"id"`
	StockID     int    `db:"stock_id"`
	IndexID     int    `db:"index_id"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	Stock       *Stock
	MarketIndex *MarketIndex
}
