package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	ClientId     string
	ClientSecret string
	Email        string
	Password     string
}

func NewConfig() *config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	return &config{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Email:        os.Getenv("EMAIL"),
		Password:     os.Getenv("PASSWORD"),
	}
}
