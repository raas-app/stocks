package scrapper

type CompanyScrapper interface {
	GetCompany(symbol string) (*Company, error)
	GetFinancials(symbol string) ([]*CompanyReport, error)
}
type Companies struct {
	count int64
	data  []Company
}

type Company struct {
	Name          string
	Symbol        string
	Sector        Sector
	Profile       CompanyProfile
	EquityProfile CompanyEquityProfile
	Announcements map[ReportType][]CompanyReport
	Financials    []*CompanyReport
}

type Sector struct {
	Name string
}
type CompanyProfile struct {
	Description   string
	Website       string
	Address       string
	FiscalYearEnd string
	Executives    []Executive
}

type Executive struct {
	Name string
	Role string
}

type CompanyEquityProfile struct {
	MarketCapitalization float64
	Shares               float64
	FreeFloat            float64
	BookValue            float64
	FaceValue            float64
}

type ReportType string

const (
	Financial    ReportType = "Financial"
	BoardMeeting ReportType = "BoardMeeting"
	Other        ReportType = "Other"
)

const (
	Quaterly ReportType = "Quaterly"
	Annual   ReportType = "Annual"
)

type CompanyReport struct {
	Date  int64
	Title string
	URL   string
	Type  ReportType
}
