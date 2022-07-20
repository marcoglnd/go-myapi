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

type BitcoinResponse struct {
	Bitcoin struct {
		Usd int `json:"usd"`
	} `json:"bitcoin"`
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

		resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
		if err != nil {
			ctx.JSON(http.StatusPartialContent, nil)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var bitResponse BitcoinResponse
		e := json.Unmarshal(body, &bitResponse)
		if e != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, bitResponse)
	}
}

// func MainController() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if ctx.Query("data") != "" {
// 			ReturnData()
// 		}
// 	}
// }
