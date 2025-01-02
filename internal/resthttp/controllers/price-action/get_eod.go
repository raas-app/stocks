package priceaction

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	priceactiondto "github.com/raas-app/stocks/internal/resthttp/dto/priceactionDto"
)

func (pa *PriceActionHandler) GetEndOfDay(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var symbol = chi.URLParam(r, "symbol")

	priceActions := pa.usecase.GetEodPriceAction(ctx, symbol)
	pa.responder.Ok(w, priceactiondto.NewEndOfDayResponse(priceActions))

}
