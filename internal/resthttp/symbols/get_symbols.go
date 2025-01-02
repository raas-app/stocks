package symbols

import (
	"context"
	"net/http"
	"time"

	symbolsdto "github.com/raas-app/stocks/internal/resthttp/dto/symbolsDto"
)

func (s *SymbolsHandler) GetSymbols(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	symbols := s.stockFetcherHandler.GetSymbols(ctx)
	s.responder.Ok(w, symbolsdto.NewGetSymbolsResponse(symbols))
}
