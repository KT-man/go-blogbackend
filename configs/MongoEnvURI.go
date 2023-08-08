package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func MongoEnvURI () string {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("dotenv did not load", envErr)
	}
	
	if envErr != nil {
		log.Fatal("Fatal error occurred")
	 }

	 return os.Getenv("MONGOURI")
}