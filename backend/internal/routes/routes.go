package routes

import (
	"fmt"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/MatheusHenrique129/application-in-go/internal/controller"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	config         *config.Config
	logger         *util.Logger
	userController *controller.UserController
	feedController *controller.FeedController
}

func (r *Routes) CreateRouter() *gin.Engine {
	r.logger.InfofWithoutContext("Mapping routes")

	// initialize gin routing engine
	router := gin.Default()

	// Endpoints
	v1Group := router.RouterGroup.Group(getVersionedApiPrefix(consts.V1))
	{
		// :: Users Endpoints
		v1Group.GET("/user/:user_id", r.userController.FindUser)
		v1Group.GET("/user", r.feedController.ListAllProducts)

	}

	// :: Swagger Routes
	router.GET("/swagger/*any", config.Swagger())

	return router
}

func getVersionedApiPrefix(version string) string {
	return fmt.Sprintf("%s/%s", consts.ApiRoutePrefix, version)
}

func NewRoutes(config *config.Config, userController *controller.UserController, feedController *controller.FeedController) *Routes {
	routes := &Routes{
		config:         config,
		userController: userController,
		feedController: feedController,
		logger:         util.NewLogger("Routes"),
	}
	return routes
}
