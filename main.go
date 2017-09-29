package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/acquirer"
	"github.com/ingresse/payment/middleware"
)

func main() {
	server := gin.New()

	// General Middlewares
	server.Use(gin.Logger())
	server.Use(middleware.Cors())
	//server.Use(middleware.Recovery())

	// Routes
	acquirer.Routes(server)

	// Run server
	server.Run()
}
