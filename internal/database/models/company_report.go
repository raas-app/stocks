package models

type CompanyReport struct {
	ID        int        `db:"id"`
	Title     string     `db:"title"`
	URL       string     `db:"url"`
	Type      ReportType `db:"type"`
	StockID   int        `db:"stock_id"`
	CreatedAt string     `db:"created_at"`
	UpdatedAt string     `db:"updated_at"`
	Stock     *Stock
}

type ReportType string

const (
	AnnualReport    ReportType = "Annual"
	QuarterlyReport ReportType = "Quarterly"
	BoardMeeting    ReportType = "BoardMeeting"
	Financial       ReportType = "Financial"
	Others          ReportType = "Other"
)
