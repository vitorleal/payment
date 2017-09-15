package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/acquirer"
	"github.com/ingresse/payment/middleware"
)

func main() {
	server := gin.Default()

	// General Middlewares
	server.Use(middleware.Cors())

	// Controllers
	acq := acquirer.NewController()

	// Define Routes
	pay := server.Group("/")
	{
		pay.POST("/", middleware.ValidatePayment(), middleware.Antifraud(), acq.Pay)
		pay.GET("/:id", acq.Get)
		pay.PUT("/:id/capture", acq.Capture)
	}

	server.Run()
}
