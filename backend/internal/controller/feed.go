package controller

import (
	"log"
	"net/http"

	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/gin-gonic/gin"
)

type FeedController struct {
	logger *util.Logger
}

func (p *FeedController) ListAllProducts(c *gin.Context) {

	result := gin.H{
		"value": "ok",
	}

	log.Print("Successful in getting the products!")
	c.JSON(http.StatusOK, result)
}

func NewProductsController() *FeedController {
	controller := &FeedController{
		logger: util.NewLogger("Feed Controller"),
	}
	return controller
}
