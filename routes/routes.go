package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/models"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/todos", models.GetAllUsers)
	router.POST("/todo", models.CreateUser)
	router.GET("/todo/:todoId", models.GetSingleUser)
	router.PUT("/todo/:todoId", models.EditUser)
	router.DELETE("/todo/:todoId", models.DeleteUser)
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
