package controller

import (
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
		id := ctx.Param("id")
		resp, err := w.webapp.GetCrypto(id)
		if err != nil {
			ctx.JSON(http.StatusPartialContent, resp)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func (w WebappController) GetRandomCrypto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := w.webapp.GetRandomCrypto()
		ctx.JSON(http.StatusOK, resp)
	}
}
