package priceaction

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	priceactiondto "github.com/raas-app/stocks/internal/resthttp/dto/priceactionDto"
)

func (pa *PriceActionHandler) GetIntraday(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var symbol = chi.URLParam(r, "symbol")
	priceActions := pa.usecase.GetIntradayPriceAction(ctx, symbol)
	pa.responder.Ok(w, priceactiondto.NewIntradayResponse(priceActions))

}
