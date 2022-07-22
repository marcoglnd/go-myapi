package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/marcoglnd/go-myapi/internal/webapp/domain"
)

type webappService struct{}

func NewWebappService() *webappService {
	return &webappService{}
}

func (w webappService) GetCryptoById(url string) (domain.CryptoResponse, error) {
	var cryptoResponse domain.CryptoResponse
	resp, err := http.Get(url)
	if err != nil {
		cryptoResponse.Partial = true
		return cryptoResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	e := json.Unmarshal(body, &cryptoResponse)
	if e != nil {
		log.Fatal(e)
	}

	return cryptoResponse, nil
}

func (w webappService) GetCrypto(id string) (domain.CryptoResponse, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
	resp, err := w.GetCryptoById(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (w webappService) GetCryptoChannel(id string, ch chan<- domain.CryptoResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
	resp, _ := w.GetCryptoById(url)
	ch <- resp
}

func (w webappService) GetRandomCrypto() []domain.CryptoResponse {
	var wg sync.WaitGroup
	var resp []domain.CryptoResponse
	wg.Add(3)
	ch := make(chan domain.CryptoResponse, 3)
	go w.GetCryptoChannel("bitcoin", ch, &wg)
	go w.GetCryptoChannel("ethereum", ch, &wg)
	go w.GetCryptoChannel("solana", ch, &wg)
	wg.Wait()
	close(ch)
	for i := 0; i < 3; i++ {
		resp = append(resp, <-ch)
	}
	return resp
}
