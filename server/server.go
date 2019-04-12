package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorleal/payment/gateway"
	"github.com/vitorleal/payment/gateway/acquirer"
	"github.com/vitorleal/payment/server/middleware"
)

// Start will configure and start the webserver
func Start() {
	server := gin.New()

	// Register Validations
	gateway.RegisterValidations()

	// General Middlewares
	server.Use(gin.Logger())
	server.Use(middleware.Cors(), middleware.Jwt())
	//server.Use(middleware.Recovery())

	// Routes
	acquirer.Routes(server)

	// Run server
	server.Run()
}
