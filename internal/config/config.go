package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required" envDefault:"3306"`
	DBName     string `env:"DB_NAME,required"`
	ServerPort string `env:"SERVER_PORT,required" envDefault:"8080"`
}

var AppConfig Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env")
	}

	err = env.Parse(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("configuration loaded successfully")
}
