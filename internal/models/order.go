package models

import "time"

type Order struct {
	ID           string
	FacilityCode string
	Status       string
	CreatedAt    time.Time
}