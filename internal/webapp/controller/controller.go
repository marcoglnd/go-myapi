package controller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/go-myapi/internal/webapp/domain"
)

type WebappController struct {
	webapp domain.WebappService
}

type DataResponse struct {
	Data string `json:"data"`
}

func NewWebappController(webapp domain.WebappService) *WebappController {
	return &WebappController{
		webapp: webapp,
	}
}

func (w WebappController) GetData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := DataResponse{
			Data: ctx.Query("data"),
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (w WebappController) GetCrypto(id string) (domain.CryptoResponse, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
	resp, err := w.webapp.GetCryptoById(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (w WebappController) GetCryptoById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		resp, err := w.GetCrypto(id)
		if err != nil {
			ctx.JSON(http.StatusPartialContent, resp)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func (w WebappController) GetCryptoChannel(id string, ch chan<- domain.CryptoResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
	resp, _ := w.webapp.GetCryptoById(url)
	ch <- resp
}

func (w WebappController) GetRandomCrypto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		ctx.JSON(http.StatusOK, resp)
	}
}
