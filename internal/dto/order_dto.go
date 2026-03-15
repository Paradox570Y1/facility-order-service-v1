package dto

import "time"

type CreateOrderRequest struct {
	ID           string `json:"id" binding:"required"`
	FacilityCode string `json:"facility_code" binding:"required"`
}

type OrderResponse struct {
	ID           string    `json:"id"`
	FacilityCode string    `json:"facility_code"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type OrderCreatedMessage struct {
	OrderID      string `json:"order_id"`
	FacilityCode string `json:"facility_code"`
}
