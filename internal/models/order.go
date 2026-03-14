package models

import "time"

type Order struct {
	ID            string    `json:"id"`
	FacilityCode  string    `json:"facility_code"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}