package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/go-myapi/internal/webapp/domain"
)

type WebappController struct {
	webapp domain.WebappService
}

func NewWebappController(webapp domain.WebappService) *WebappController {
	return &WebappController{
		webapp: webapp,
	}
}

func (w WebappController) GetData(ctx *gin.Context) {
	response := domain.DataResponse{
		Data: ctx.Query("data"),
	}
	ctx.JSON(http.StatusOK, response)
}

func (w WebappController) GetCryptoById(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := w.webapp.GetCryptoById(id)
	if err != nil {
		ctx.JSON(http.StatusPartialContent, resp)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (w WebappController) GetRandomCrypto(ctx *gin.Context) {
	resp := w.webapp.GetRandomCrypto()
	for _, crypto := range resp {
		if crypto.Partial {
			ctx.JSON(http.StatusPartialContent, resp)
			return
		}
	}
	ctx.JSON(http.StatusOK, resp)
}
