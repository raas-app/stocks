package models

type StockPriceHistory struct {
	ID         int     `db:"id"`
	StockID    int     `db:"stock_id"`
	Price      float64 `db:"price"`
	Volume     int     `db:"volume"`
	OpenPrice  float64 `db:"open_price"`
	ClosePrice float64 `db:"close_price"`
	HighPrice  float64 `db:"high_price"`
	LowPrice   float64 `db:"low_price"`
	Timestamp  string  `db:"timestamp"`
	CreatedAt  string  `db:"created_at"`
	UpdatedAt  string  `db:"updated_at"`
	Stock      *Stock
}
