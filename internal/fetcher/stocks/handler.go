package fetcher

type Stock struct {
	Symbol string
	Logo   string
	Name   string
	Sector Sector
	IsEtf  bool
	IsDebt bool
}

type Sector struct {
	ID   int64
	Name string
}

func StockHandler() {

}
