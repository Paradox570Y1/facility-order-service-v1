package routes

import (
	"database/sql"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/handlers"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/repository"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	facilityRepo := repository.NewFacilityRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	facilityService := services.NewFacilityService(facilityRepo)
	orderService := services.NewOrderService(orderRepo, facilityRepo)

	facilityHandler := handlers.NewFacilityHandler(facilityService)
	orderHandler := handlers.NewOrderHandler(orderService)

	facilities := router.Group("/facilities")
	{
		facilities.GET("", facilityHandler.GetAll)
		facilities.GET("/:code", facilityHandler.GetByCode)
		facilities.POST("", facilityHandler.Create)
	}

	orders := router.Group("/orders")
	{
		orders.GET("", orderHandler.GetAll)
		orders.GET("/:id", orderHandler.GetByID)
		orders.POST("", orderHandler.Create)
	}
}
