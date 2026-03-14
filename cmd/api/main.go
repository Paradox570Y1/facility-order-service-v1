package main

import (
	"github.com/Paradox570Y1/facility-order-service-v1/internal/config"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/db"
)

func main() {
	config.Load()
	db.Connect()
}