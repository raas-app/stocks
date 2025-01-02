package priceactiondto

import (
	"time"

	raas "github.com/raas-app/stocks"
)

type GetIntradayResponse struct {
	Count int            `json:"count"`
	Data  []IntraDayData `json:"data"`
}

type IntraDayData struct {
	Price  float64   `json:"price"`
	Volume float64   `json:"volume"`
	Time   TimeField `json:"time"`
}

func NewIntradayResponse(data []raas.IntradayPriceAction) *GetIntradayResponse {
	var result = make([]IntraDayData, len(data))
	for index, priceAction := range data {
		t := time.Unix(int64(priceAction.Time), 0)
		timeOnly := t.Format(time.TimeOnly)
		date := t.Format(time.DateOnly)
		result[index] = IntraDayData{
			Price:  priceAction.Price,
			Volume: priceAction.Volume,
			Time: TimeField{
				TimeOnly: timeOnly,
				DateOnly: date,
				Epoch:    int64(priceAction.Time),
			},
		}
	}
	return &GetIntradayResponse{
		Count: len(data),
		Data:  result,
	}
}
