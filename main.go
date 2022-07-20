package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/go-myapi/controller"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/myapi", controller.ReturnData())

	r.Run()
}
