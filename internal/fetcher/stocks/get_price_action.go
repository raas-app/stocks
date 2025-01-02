package stocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	raas "github.com/raas-app/stocks"
	"go.uber.org/zap"
)

func (h *stockHandler) GetIntradayPriceAction(ctx context.Context, symbol string) []raas.IntradayPriceAction {
	url := fmt.Sprintf("%s%s%s", h.Config.Market.PSX.BaseURL, h.Config.Market.PSX.TimeseriesURL.Intraday, symbol)
	var intradayPriceActions []raas.IntradayPriceAction
	err := h.fetchAndParse(url, &intradayPriceActions)
	if err != nil {
		h.Logger.Error("Failed to fetch intraday price action", zap.String("symbol", symbol), zap.Error(err))
		return nil
	}
	return intradayPriceActions
}

func (h *stockHandler) GetEodPriceAction(ctx context.Context, symbol string) []raas.EodPriceAction {
	url := fmt.Sprintf("%s%s%s", h.Config.Market.PSX.BaseURL, h.Config.Market.PSX.TimeseriesURL.EOD, symbol)
	var eodPriceActions []raas.EodPriceAction
	err := h.fetchAndParse(url, &eodPriceActions)
	if err != nil {
		h.Logger.Error("Failed to fetch EOD price action", zap.String("symbol", symbol), zap.Error(err))
		return nil
	}
	return eodPriceActions
}

// fetchAndParse handles the HTTP request and parses the JSON response into the given target.
func (h *stockHandler) fetchAndParse(url string, target interface{}) error {
	// Make HTTP Request
	resp, err := h.makeRequest(url)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			h.Logger.Error("Failed to close response body", zap.Error(cerr))
		}
	}()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return h.parseResponse(body, target)
}

// makeRequest creates and executes an HTTP GET request.
func (*stockHandler) makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	return resp, nil
}

// parseResponse unmarshals the JSON response into the target struct.
func (h *stockHandler) parseResponse(body []byte, target interface{}) error {
	var response PriceActionResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	switch target := target.(type) {
	case *[]raas.IntradayPriceAction:
		*target = h.parseIntradayPriceActions(response.Data)
	case *[]raas.EodPriceAction:
		*target = h.parseEodPriceActions(response.Data)
	default:
		return fmt.Errorf("unsupported target type")
	}

	return nil
}

// parseIntradayPriceActions converts raw response data into IntradayPriceAction slice.
func (*stockHandler) parseIntradayPriceActions(data [][]float64) []raas.IntradayPriceAction {
	priceActions := make([]raas.IntradayPriceAction, len(data))
	for i, priceAction := range data {
		priceActions[i] = raas.IntradayPriceAction{
			Time:   priceAction[0],
			Price:  priceAction[1],
			Volume: priceAction[2],
		}
	}
	return priceActions
}

// parseEodPriceActions converts raw response data into EodPriceAction slice.
func (*stockHandler) parseEodPriceActions(data [][]float64) []raas.EodPriceAction {
	priceActions := make([]raas.EodPriceAction, len(data))
	for i, priceAction := range data {
		priceActions[i] = raas.EodPriceAction{
			Time:   priceAction[0],
			Close:  priceAction[1],
			Volume: priceAction[2],
			Open:   priceAction[3],
		}
	}
	return priceActions
}
