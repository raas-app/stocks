package models

type Sector struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Stocks []*Stock

	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
