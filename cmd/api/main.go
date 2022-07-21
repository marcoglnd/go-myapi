package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/marcoglnd/go-myapi/internal/webapp/controller"
	"github.com/marcoglnd/go-myapi/internal/webapp/service"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	webappService := service.NewWebappService()
	webappController := controller.NewWebappController(webappService)
	r := gin.Default()
	r.GET("/myapi", webappController.GetData())
	r.GET("/crypto/:id", webappController.GetCryptoById())
	// r.GET("/randomcrypto", webappController.GetRandomCrypto())

	r.Run()
}
