package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/internal/domain/controller"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", welcome)
	router.POST("/user", controller.CreateUser())
	router.GET("/user/:userId", controller.GetAUser())
	router.PUT("/user/:userId", controller.EditAUser())
	router.DELETE("/user/:userId", controller.DeleteAUser())
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
