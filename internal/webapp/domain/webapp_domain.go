package domain

import "sync"

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
	GetCryptoById(url string) (CryptoResponse, error)
	GetCrypto(id string) (CryptoResponse, error)
	GetCryptoChannel(id string, ch chan<- CryptoResponse, wg *sync.WaitGroup)
	GetRandomCrypto() []CryptoResponse
}
