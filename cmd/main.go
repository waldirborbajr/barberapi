package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/config"
	"github.com/waldirborbajr/barberapi/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {}

func main() {
	var DB *mongo.Client = config.ConnectDB()

	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":9090"))
}
