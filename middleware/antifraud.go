package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/antifraud/siftscience"
	"github.com/ingresse/payment/payment"
)

// Antifraud middleware
func Antifraud() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.MustGet("body").(payment.Payment)

		if body.Antifraud != "siftscience" {
			c.Next()
			return
		}

		sift := siftscience.New(siftscience.Sandbox)
		score, err := sift.GetScore(body.Customer.Id)

		if err != nil {
			c.Error(err).SetMeta("Middleware.Antifraud.Sift.Api.Error")

			c.AbortWithStatusJSON(400, gin.H{
				"error": err,
			})
			return
		}

		if score == nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Error getting score of user " + body.Customer.Id,
			})
			return
		}

		if score.Scores.PaymentAbuse.Score >= 0.4 {
			c.AbortWithStatusJSON(400, gin.H{
				"error":     "User with score over the limit",
				"antifraud": score,
			})

			return
		}

		c.Next()
	}
}
