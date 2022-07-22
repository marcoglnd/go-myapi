package controller

import (
	"fmt"
	"net/http"

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

func (w WebappController) GetRandomCrypto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var resp []domain.CryptoResponse
		resp1, err1 := w.GetCrypto("bitcoin")
		resp2, err2 := w.GetCrypto("ethereum")
		resp3, err3 := w.GetCrypto("solana")
		resp = append(resp, resp1, resp2, resp3)
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusPartialContent, resp)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
