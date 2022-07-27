package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/go-myapi/internal/webapp/controller"
	"github.com/marcoglnd/go-myapi/internal/webapp/service"
)

func webappRouter(superRouter *gin.RouterGroup) {
	webappService := service.NewWebappService()
	webappController := controller.NewWebappController(webappService)

	router := superRouter.Group("crypto")
	{
		router.GET("/myapi", webappController.GetData)
		router.GET("/:id", webappController.GetCryptoById)
		router.GET("/randomcrypto", webappController.GetRandomCrypto)
	}
}
