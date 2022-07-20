package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data string `json:"data"`
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

		// resp, err := http.Get("")
	}
}

// func MainController() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if ctx.Query("data") != "" {
// 			ReturnData()
// 		}
// 	}
// }
