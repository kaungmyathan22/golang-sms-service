package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envAccountSID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file.")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAccountAuthToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file.")
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func envAccountServiceID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file.")
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}
