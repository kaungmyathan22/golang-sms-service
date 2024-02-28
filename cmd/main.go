package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-otp-sms-service/api"
)

func main() {
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8000")
}
