package db

import (
	"log"
	"fmt"
	"time"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db;
	log.Println("Database connected successfully")
}