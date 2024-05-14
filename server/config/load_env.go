package config

import (
	"log"

	"github.com/lpernett/godotenv"
)

func GetEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("could not find .env file")
	}
}
