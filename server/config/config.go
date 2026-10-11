package config




type Config struct {
	Port string 
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found; using system environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	return Config{
		Port: port,
	}
}