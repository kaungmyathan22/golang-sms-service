package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-otp-sms-service/api"
)

func main() {
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	//routes
	app.Routes()
	log.Println("Server is running at http://localhost:8000")
	router.Run(":8000")
}
