package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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
