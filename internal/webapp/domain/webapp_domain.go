package domain

import "sync"

type DataResponse struct {
	Data string `json:"data"`
}

type CurrentPrice struct {
	Usd float64 `json:"usd,omitempty"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price,omitempty"`
}

type CryptoResponse struct {
	ID         string     `json:"id"`
	Symbol     string     `json:"symbol,omitempty"`
	MarketData MarketData `json:"market_data,omitempty"`
	Partial    bool       `json:"partial"`
}

type WebappService interface {
	GetCryptoById(id string) (CryptoResponse, error)
	GetCryptoUrl(id string) string
	GetCryptoChannel(id string, ch chan<- CryptoResponse, wg *sync.WaitGroup)
	GetRandomCrypto() []CryptoResponse
}
