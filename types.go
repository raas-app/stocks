package raas

type IntradayPriceAction struct {
	Price  float64
	Volume float64
	Time   float64
}

type EodPriceAction struct {
	Time   float64
	Open   float64
	Close  float64
	Volume float64
}
type Stock struct {
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	SectorName string `json:"sectorName"`
	IsETF      bool   `json:"isETF"`
	IsDebt     bool   `json:"isDebt"`
}
