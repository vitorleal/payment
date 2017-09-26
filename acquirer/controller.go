package acquirer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/acquirer/cielo"
	"github.com/ingresse/payment/acquirer/stone"
	"github.com/ingresse/payment/gateway"
)

// Acquirer controller
type AcquirerController struct {
	Cielo *cielo.Client
	Stone *stone.Client
}

// Init the controller
func NewController() *AcquirerController {
	// Instantiate a new Cielo client
	cieloClient := cielo.New(cielo.Merchant{
		Id:  "7df78036-fe0a-4909-9315-933ccb3ab5cd",
		Key: "QDWVJXLMPKWFXWDMWLWAGUDHBHMVQVHJOWLQYZGQ",
	}, cielo.Sandbox)

	// Instantiate a new Stone client
	stoneClient := stone.New(stone.Merchant{
		Key: "f2a1f485-cfd4-49f5-8862-0ebc438ae923",
	}, stone.Production)

	return &AcquirerController{cieloClient, stoneClient}
}

// Create a new payment
func (controller *AcquirerController) Pay(c *gin.Context) {
	payment := c.MustGet("body").(gateway.Payment)

	// Handle cielo payment
	if payment.IsAcquirer(cielo.Name) {
		c.JSON(http.StatusOK, gin.H{
			"acquirer": cielo.Name,
			"sale":     nil,
		})

		return
	}

	// handle stone payment
	if payment.IsAcquirer(stone.Name) {
		sale := stone.FormatSale(&payment)
		response, err := controller.Stone.NewSale(sale)

		// If error exist
		if err.ErrorReport != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err,
			})
			return
		}

		// Return the transaction
		c.JSON(http.StatusOK, gin.H{
			"acquirer": stone.Name,
			"sale":     response,
		})

		return
	}

	// acquirer not implemented
	c.AbortWithStatusJSON(400, gin.H{
		"error": "Invalid acquirer " + payment.Acquirer,
	})
}

// Capture an existing transaction
func (controller *AcquirerController) Capture(c *gin.Context) {
	response, err := controller.Cielo.CaptureSale(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"acquirer": "cielo",
		"sale":     response,
	})
}

// Get data of an existing transaction
func (controller *AcquirerController) Get(c *gin.Context) {
	id := c.Param("id")

	response, err := controller.Stone.GetSale(id)

	// If error exist
	if err.ErrorReport != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err,
		})
		return
	}

	// Return the transaction
	c.JSON(http.StatusOK, gin.H{
		"acquirer": stone.Name,
		"sale":     response,
	})

	return
}
