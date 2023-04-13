package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/config"
	"github.com/waldirborbajr/barberapi/routes"
)

func init() {}

func main() {
	router := gin.Default()

	config.ConnectDB()

	routes.UserRoutes(router)

	log.Fatal(router.Run(":9090"))
}
