package models

type MarketIndex struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CountryID int    `db:"country_id"`
	Country   *Country
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
