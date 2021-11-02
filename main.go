package main

import (
	"log"

	"github.com/CRYBOII/kill-ecommerce/configs"
	"github.com/CRYBOII/kill-ecommerce/src/routes"
	valid "github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	valid.SetFieldsRequiredByDefault(true)
	if configs.CheckConnection() == 0 {
		log.Fatal("Connection failed")
		return
	}
	routes.Routes()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
