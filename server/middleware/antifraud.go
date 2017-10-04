package middleware

import (
	"github.com/gin-gonic/gin"
	g "github.com/ingresse/payment/gateway"
	"github.com/ingresse/payment/gateway/antifraud/siftscience"
)

// Antifraud is a middleware to validate customer
// cleaerance in the antifrauds services
func Antifraud() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.MustGet("body").(g.Payment)

		// If not antifraud declared in the body continue
		if body.Antifraud == nil {
			c.Next()
			return
		}

		for _, a := range body.Antifraud {
			// If siftscience
			if a.Name == siftscience.Name {
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

				if score.Scores.PaymentAbuse.Score >= a.Score {
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
