package routes

import (
	"fmt"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/constants"
	"github.com/MatheusHenrique129/application-in-go/internal/controller/v1"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	config *config.Config
	logger *util.Logger
}

func (r *Routes) CreateRouter() *gin.Engine {
	r.logger.InfofWithoutContext("Mapping routes")

	// initialize gin routing engine
	router := gin.Default()

	// V1 endpoints
	v1Group := router.RouterGroup.Group(getVersionedApiPrefix(constants.V1))
	{
		// :: Products Endpoints
		v1Group.GET("/product", v1.Products)

	}

	// :: Swagger Routes
	router.GET("/swagger/*any", config.Swagger())

	return router
}

func getVersionedApiPrefix(version string) string {
	return fmt.Sprintf("%s/%s", constants.ApiRoutePrefix, version)
}

func NewRoutes(config *config.Config) *Routes {
	routes := &Routes{
		config: config,
		logger: util.NewLogger("Routes"),
	}
	return routes
}
