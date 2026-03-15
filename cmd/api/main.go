package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Paradox570Y1/facility-order-service-v1/internal/config"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/db"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/kafka"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/routes"
)

func main() {
	config.Load()
	db.Connect()
	db.RunMigrations()

	// Start Kafka consumer in background
	go kafka.ConsumerWorker(context.Background(), config.AppConfig.KafkaBrokers)

	router := gin.Default()
	routes.RegisterRoutes(router, db.DB, config.AppConfig.KafkaBrokers)

	log.Printf("Server starting on port %s", config.AppConfig.ServerPort)
	router.Run(":" + config.AppConfig.ServerPort)
}
