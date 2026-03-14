package handlers

import (
	"net/http"

	"github.com/Paradox570Y1/facility-order-service-v1/internal/dto"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	orders, err := h.orderService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderService.GetByID(c.Request.Context(), id)
	if err != nil {
		if services.IsOrderNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req dto.CreateOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.orderService.Create(c.Request.Context(), req)
	if err != nil {
		if services.IsInvalidID(err) || services.IsFacilityDoesNotExist(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}
