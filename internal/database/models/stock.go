package models

import "encoding/json"

type Stock struct {
	ID             int             `db:"id"`
	Symbol         string          `db:"symbol"`
	Name           string          `db:"name"`
	Price          float64         `db:"price"`
	MarketCap      float64         `db:"market_cap"`
	Shares         int             `db:"shares"`
	FreeFloat      int             `db:"free_float"`
	BookValue      float64         `db:"book_value"`
	FaceValue      float64         `db:"face_value"`
	CompanyID      int             `db:"company_id"`
	SectorID       int             `db:"sector_id"`
	CompanyProfile json.RawMessage `db:"company_profile"`
	Sector         *Sector
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}
