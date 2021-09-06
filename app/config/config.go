package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var USER_URL string
var GATEWAY_URL string
var PAYMENT_URL string
var MERCHANT_URL string

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	USER_URL = os.Getenv("USER_URL") + "v1/"
	GATEWAY_URL = os.Getenv("GATEWAY_URL") + "v1/"
	PAYMENT_URL = os.Getenv("PAYMENT_URL") + "v1/"
	MERCHANT_URL = os.Getenv("MERCHANT_URL") + "v1/"
}
