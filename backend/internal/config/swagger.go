package config

import (
	"github.com/MatheusHenrique129/application-in-go/internal/constants"
	"github.com/MatheusHenrique129/application-in-go/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Swagger creates the Swagger handler for Gin.
func Swagger() gin.HandlerFunc {
	docs.SwaggerInfo.Version = constants.V1
	docs.SwaggerInfo.Schemes = []string{"https"}

	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
