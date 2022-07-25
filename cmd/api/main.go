package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/go-myapi/cmd/api/routes"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	PATH := "api/v1"
	router := gin.Default()
	routerGroup := router.Group(PATH)
	routes.AddRoutes(routerGroup)
	err := router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
