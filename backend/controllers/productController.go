package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Products(c *gin.Context) {

	result := gin.H{
		"value": "ok",
	}

	log.Print("Successful in getting the products!")
	c.JSON(http.StatusOK, result)
}
