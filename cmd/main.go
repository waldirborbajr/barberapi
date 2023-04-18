package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/waldirborbajr/barberapi/config"
	"github.com/waldirborbajr/barberapi/internal/infra/db"
	"github.com/waldirborbajr/barberapi/routes"
)

func init() {}

func main() {
	router := gin.Default()

	configs, err := config.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	fmt.Println(configs)

	db.ConnectDB()

	routes.UserRoutes(router)

	log.Fatal(router.Run(":9090"))
}
