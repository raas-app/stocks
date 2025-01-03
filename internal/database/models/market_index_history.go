package models

type MarketIndexHistory struct {
	ID          int     `db:"id"`
	IndexID     int     `db:"index_id"`
	Date        string  `db:"date"`
	OpenPrice   float64 `db:"open_price"`
	ClosePrice  float64 `db:"close_price"`
	HighPrice   float64 `db:"high_price"`
	LowPrice    float64 `db:"low_price"`
	Volume      int     `db:"volume"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
	MarketIndex *MarketIndex
}
