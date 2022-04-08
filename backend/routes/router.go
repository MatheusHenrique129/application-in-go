package routes

import (
	"github.com/MatheusHenrique129/application-in-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.Products)
		}
	}

	return router
}
