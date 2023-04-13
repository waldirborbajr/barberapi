package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/models"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/users", models.GetAllUsers)
	router.POST("/user", models.CreateUser)
	router.GET("/user/:userId", models.GetSingleUser)
	router.PUT("/user/:userId", models.EditUser)
	router.DELETE("/user/:userId", models.DeleteUser)
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
