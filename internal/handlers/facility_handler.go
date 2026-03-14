package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/dto"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/services"
)

type FacilityHandler struct {
	facilityService services.FacilityService
}

func NewFacilityHandler(facilityService services.FacilityService) *FacilityHandler {
	return &FacilityHandler{
		facilityService: facilityService,
	}
}

func (h *FacilityHandler) GetAll(c *gin.Context) {
	facilities, err := h.facilityService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, facilities)
}

func (h *FacilityHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")

	facility, err := h.facilityService.GetByCode(c.Request.Context(), code)
	if err != nil {
		if err.Error() == "facility not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, facility)
}

func (h *FacilityHandler) Create(c *gin.Context) {
	var req dto.CreateFacilityRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.facilityService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "facility created successfully"})
}
