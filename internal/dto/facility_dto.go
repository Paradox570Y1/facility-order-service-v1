package dto

type CreateFacilityRequest struct {
	Code    string `json:"code" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type FacilityResponse struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
}