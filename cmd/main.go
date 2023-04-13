package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Barber Schedule up and running!!!"})
	})

	router.Run(":9090")
}
