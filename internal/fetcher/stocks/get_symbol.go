package stocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (h *stockHandler) GetSymbols() []Stock {
	// Construct the URL
	url := fmt.Sprintf("%s%s", h.Config.Market.PSX.BaseURL, h.Config.Market.PSX.ScraperURL.Symbols)

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		h.Logger.Error("Failed to create request for symbols", zap.Error(err))
		return nil
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		h.Logger.Error("HTTP request failed", zap.String("url", url), zap.Error(err))
		return nil
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			h.Logger.Error("Failed to close response body", zap.Error(cerr))
		}
	}()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		h.Logger.Error("Unexpected status code from symbols endpoint",
			zap.String("url", url), zap.Int("statusCode", resp.StatusCode))
		return nil
	}

	// Read and parse response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		h.Logger.Error("Failed to read response body", zap.Error(err))
		return nil
	}

	var symbols []Stock
	if err := json.Unmarshal(body, &symbols); err != nil {
		h.Logger.Error("Failed to parse JSON response",
			zap.String("url", url), zap.ByteString("responseBody", body), zap.Error(err))
		return nil
	}

	// Log and return parsed symbols
	h.Logger.Info("Fetched symbols successfully", zap.Int("count", len(symbols)))
	return symbols
}
