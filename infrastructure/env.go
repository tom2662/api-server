package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

//load env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}
