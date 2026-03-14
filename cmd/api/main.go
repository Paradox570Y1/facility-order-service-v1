package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Paradox570Y1/facility-order-service-v1/internal/config"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/db"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/routes"
)

func main() {
	config.Load()
	db.Connect()
	db.RunMigrations()
	router := gin.Default()
	routes.RegisterRoutes(router, db.DB)
	router.Run(":" + config.AppConfig.ServerPort)
}
