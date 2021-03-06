package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/vitorleal/payment/gateway"
)

// Validate payment body before procedd
func ValidatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body gateway.Payment

		// Validate the JSON body
		bindError := c.ShouldBindWith(&body, binding.JSON)

		// Return error if exisst and abort request
		if bindError != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": bindError,
			})

			c.Error(bindError).SetMeta("Middleware.ValidatePayment.Invalid")
			return
		}

		// Set the content in context to be used in the next handlers
		c.Set("body", body)
		c.Next()
	}
}
