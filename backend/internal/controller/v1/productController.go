package v1

import (
	"log"
	"net/http"

	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	logger *util.Logger
}

func Products(c *gin.Context) {

	result := gin.H{
		"value": "ok",
	}

	log.Print("Successful in getting the products!")
	c.JSON(http.StatusOK, result)
}

func NewProductsController() *ProductsController {
	controller := &ProductsController{
		logger: util.NewLogger("Products Controller"),
	}
	return controller
}
