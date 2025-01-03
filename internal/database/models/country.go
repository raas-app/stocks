package models

type Country struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Code      string `db:"code"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
