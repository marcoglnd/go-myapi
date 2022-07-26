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

func (w webappService) GetCryptoUrl(id string) string {
	url := fmt.Sprintf("aaa%s", id)
	return url
}

func (w webappService) GetCryptoById(id string) (domain.CryptoResponse, error) {
	url := w.GetCryptoUrl(id)
	var cryptoResponse domain.CryptoResponse
	resp, err := http.Get(url)

	if err != nil {
		cryptoResponse.Partial = true
		return cryptoResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		log.Fatal(err)
	}

	e := json.Unmarshal(body, &cryptoResponse)
	if e != nil {
		log.Fatal(e)
	}

	return cryptoResponse, nil
}

func (w webappService) GetCryptoChannel(id string, ch chan<- domain.CryptoResponse, wg *sync.WaitGroup) {
	defer w.Recover(id, ch)
	defer wg.Done()
	resp, _ := w.GetCryptoById(id)
	fmt.Println(resp)
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

func (w webappService) Recover(id string, ch chan<- domain.CryptoResponse) {
	if r := recover(); r != nil {
		ch <- domain.CryptoResponse{
			ID:      id,
			Partial: true,
		}
	}
}
