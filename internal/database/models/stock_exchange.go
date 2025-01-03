package models

type StockExchange struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Symbol    string `db:"symbol"`
	CountryID int    `db:"country_id"`
	Country   *Country
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
