package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/myapi", func(c *gin.Context) {
		response := Response{
			Data: c.Query("data"),
		}
		c.JSON(http.StatusOK, response)
	})

	r.Run()
}
