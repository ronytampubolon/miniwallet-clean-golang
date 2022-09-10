package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ronytampubolon/miniwallet/config"
	"github.com/ronytampubolon/miniwallet/routes"
	"github.com/ronytampubolon/miniwallet/utils"
)

func main() {
	db := config.Connection()
	router := gin.Default()

	// allow cors origin
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	routes.InitRoute(db, router)
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}
