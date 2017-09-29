package acquirer

import (
	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/middleware"
)

// Acquirer routes
func Routes(server *gin.Engine) {
	controller := Controller{}

	acquirer := server.Group("/")
	{
		// Get payment data
		acquirer.GET("/:acquirer/:id", controller.Get)

		// Authorize new payment
		acquirer.POST("/", middleware.ValidatePayment(), middleware.Antifraud(), controller.Authorize)

		// Capture a payment
		acquirer.POST("/:id/capture", middleware.ValidatePayment(), controller.Capture)
	}
}
