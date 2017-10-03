package acquirer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/gateway"
)

// Acquirer controller
type Controller struct{}

// Create a new payment
func (controller *Controller) Authorize(c *gin.Context) {
	payment := c.MustGet("body").(gateway.Payment)

	acquirer := NewAcquirer(payment.Acquirer)
	response, err := acquirer.Authorize(&payment)

	// If error exist
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	// Return the transaction
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

	return
}

// Capture an existing transaction
func (controller *Controller) Capture(c *gin.Context) {
	payment := c.MustGet("body").(gateway.Payment)

	acquirer := NewAcquirer(payment.Acquirer)
	response, err := acquirer.Capture(payment.Id)

	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// Get data of an existing transaction
func (controller *Controller) Get(c *gin.Context) {
	id := c.Param("id")

	acquirer := NewAcquirer("stone")
	response, err := acquirer.Get(id)

	// If error exist
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	// Return the transaction
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

	return
}
