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
	return &AcquirerController{
		cielo.New(cielo.Merchant{
			Id:  "7df78036-fe0a-4909-9315-933ccb3ab5cd",
			Key: "QDWVJXLMPKWFXWDMWLWAGUDHBHMVQVHJOWLQYZGQ",
		}, cielo.Sandbox),
		stone.New(stone.Merchant{
			Key: "f2a1f485-cfd4-49f5-8862-0ebc438ae923",
		}, stone.Production),
	}
}

// Execute a payment
func (controller *AcquirerController) Pay(c *gin.Context) {
	payment := c.MustGet("body").(gateway.Payment)

	if payment.IsCielo() {
		c.JSON(http.StatusOK, gin.H{
			"acquirer": "cielo",
			"order":    nil,
		})

		return
	}

	if payment.IsStone() {
		cardTransaction := stone.CreditCardTransaction{
			AmountInCents: payment.Amount,
			CreditCard: &stone.CreditCard{
				CreditCardBrand:  payment.CreditCard.Brand,
				CreditCardNumber: payment.CreditCard.Number,
				HolderName:       payment.CreditCard.Holder,
				ExpMonth:         payment.CreditCard.Expiration[:2],
				ExpYear:          payment.CreditCard.Expiration[3:],
				SecurityCode:     payment.CreditCard.CVV,
			},
			Options: &stone.Options{
				PaymentMethodCode: 1,
			},
		}

		sale := stone.Sale{
			CreditCardTransactionCollection: []*stone.CreditCardTransaction{&cardTransaction},
			Order: &stone.Order{
				OrderReference: payment.Id,
			},
		}

		response, err := controller.Stone.NewSale(&sale)

		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"acquirer": "stone",
			"order":    response,
		})

		return
	}

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
		"order":    response,
	})
}

// Get data of an existing transaction
func (controller *AcquirerController) Get(c *gin.Context) {
	response, err := controller.Cielo.GetSale(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"acquirer": "cielo",
		"order":    response,
	})
}
