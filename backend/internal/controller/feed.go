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

// ListAllProducts List of roles.
// @Summary Returns a list of products.
// @Tags products
// @Produce json
// @Success 200 {object} domain.RoleListResource
// @Failure 400 {object} apierrors.apiErr
// @Failure 500 {object} apierrors.apiErr
// @Router /v1/products [get]
func (p *FeedController) ListAllProducts(c *gin.Context) {

	result := gin.H{
		"value": "ok",
	}

	log.Print("Successful in getting the products!")
	c.JSON(http.StatusOK, result)
}

func NewProductsController() *FeedController {
	controller := &FeedController{
		logger: util.NewLogger("FindUser Controller"),
	}
	return controller
}
