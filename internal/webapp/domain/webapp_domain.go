package domain

import "sync"

type DataResponse struct {
	Data string `json:"data"`
}

type CurrentPrice struct {
	Usd float64 `json:"usd"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CryptoResponse struct {
	ID         string     `json:"id"`
	Symbol     string     `json:"symbol"`
	MarketData MarketData `json:"market_data"`
	Partial    bool       `json:"partial"`
}

type WebappService interface {
	GetCryptoById(url string) (CryptoResponse, error)
	GetCryptoUrl(id string) string
	GetCryptoChannel(id string, ch chan<- CryptoResponse, wg *sync.WaitGroup)
	GetRandomCrypto() []CryptoResponse
}
