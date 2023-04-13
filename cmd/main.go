package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/config"
	"github.com/waldirborbajr/barberapi/routes"
)

func init() {}

func main() {
	config.Connect()

	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":9090"))
}
