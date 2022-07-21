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

func (w WebappController) GetCryptoById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", ctx.Param("id"))
		resp, err := w.webapp.GetCryptoById(ctx, url)
		if err != nil {
			ctx.JSON(http.StatusPartialContent, resp)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

// func (w WebappController) GetRandomCrypto() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
