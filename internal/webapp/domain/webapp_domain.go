package domain

import "context"

type CryptoResponse struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	MarketData struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
	Partial bool `json:"partial"`
}

type WebappService interface {
	GetCryptoById(ctx context.Context, url string) (CryptoResponse, error)
}
