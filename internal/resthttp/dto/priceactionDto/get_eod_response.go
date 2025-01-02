package priceactiondto

import (
	"time"

	raas "github.com/raas-app/stocks"
)

type GetEndOfDayResponse struct {
	Count int            `json:"count"`
	Data  []EndOfDayData `json:"data"`
}

type EndOfDayData struct {
	Open   float64   `json:"open"`
	Close  float64   `json:"close"`
	Volume float64   `json:"volume"`
	Time   TimeField `json:"date"`
}

type TimeField struct {
	TimeOnly string `json:"time"`
	Epoch    int64  `json:"epoch"`
	DateOnly string `json:"date"`
}

func NewEndOfDayResponse(data []raas.EodPriceAction) *GetEndOfDayResponse {
	var result = make([]EndOfDayData, len(data))
	for index, priceAction := range data {
		t := time.Unix(int64(priceAction.Time), 0)

		result[index] = EndOfDayData{
			Open:   priceAction.Open,
			Close:  priceAction.Close,
			Volume: priceAction.Volume,
			Time: TimeField{
				Epoch:    int64(priceAction.Time),
				DateOnly: t.Format(time.RFC1123),
			},
		}
	}
	return &GetEndOfDayResponse{
		Count: len(data),
		Data:  result,
	}
}
