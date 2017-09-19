package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/antifraud/siftscience"
	"github.com/ingresse/payment/gateway"
)

// Antifraud middleware
func Antifraud() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.MustGet("body").(gateway.Payment)

		if body.Antifraud == nil {
			c.Next()
			return
		}

		for _, antifraud := range body.Antifraud {
			// Validate the antifraud services
			if !antifraud.IsValid() {
				c.AbortWithStatusJSON(400, gin.H{
					"error": "Invalid antifraud option: " + antifraud,
				})
				return
			}

			// If siftscience
			if antifraud.String() == siftscience.Name {
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
			}
		}

		c.Next()
	}
}
