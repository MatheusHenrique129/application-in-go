package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {

	result := gin.H{
		"value": "ok",
	}

	c.JSON(http.StatusOK, result)
}
