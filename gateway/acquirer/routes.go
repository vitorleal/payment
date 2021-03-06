package acquirer

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorleal/payment/server/middleware"
)

// Acquirer routes
func Routes(server *gin.Engine) {
	controller := Controller{}

	acquirer := server.Group("/")
	{
		// Authorize new payment
		acquirer.POST("/", middleware.ValidatePayment(), controller.Authorize)

		// Get payment data
		acquirer.GET("/:id", controller.Get)

		// Capture a payment
		acquirer.POST("/:id/capture", middleware.ValidatePayment(), controller.Capture)
	}
}
