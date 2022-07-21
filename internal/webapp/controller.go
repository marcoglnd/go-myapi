package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data string `json:"data"`
}

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

func ReturnData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Query("data") != "" {
			response := Response{
				Data: ctx.Query("data"),
			}
			ctx.JSON(http.StatusOK, response)
			return
		}

		var cryptoResponse CryptoResponse
		resp, err := http.Get("https://api.coingecko.com/api/v3/coins/bitcoin")
		if err != nil {
			cryptoResponse.Partial = true
			ctx.JSON(http.StatusPartialContent, cryptoResponse)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		e := json.Unmarshal(body, &cryptoResponse)
		if e != nil {
			log.Fatal(e)
		}

		ctx.JSON(http.StatusOK, cryptoResponse)
	}
}
