package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/marcoglnd/go-myapi/internal/webapp"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/myapi", controller.ReturnData())

	r.Run()
}
