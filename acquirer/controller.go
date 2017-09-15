package acquirer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/acquirer/cielo"
)

// Acquirer controller
type AcquirerController struct {
	Cielo *cielo.Client
}

// Init the controller
func NewController() *AcquirerController {
	return &AcquirerController{
		cielo.New(cielo.Merchant{
			Id:  "7df78036-fe0a-4909-9315-933ccb3ab5cd",
			Key: "QDWVJXLMPKWFXWDMWLWAGUDHBHMVQVHJOWLQYZGQ",
		}, cielo.Sandbox),
	}
}

// Execute a payment
func (controller *AcquirerController) Pay(c *gin.Context) {
	//body := c.MustGet("body").(payment.Payment)

	//order := cielo.Order{
	//MerchantOrderId: body.Id,
	//Customer: &cielo.Customer{
	//Name: body.Customer.Name,
	//},
	//Payment: &cielo.Payment{
	//Type:           body.Payment.Type,
	//Amount:         body.Payment.Amount,
	//Installments:   body.Payment.Installments,
	//SoftDescriptor: body.Payment.SoftDescriptor,
	//CreditCard: &cielo.CreditCard{
	//CardNumber:     body.Payment.CreditCard.Number,
	//Holder:         body.Payment.CreditCard.Holder,
	//ExpirationDate: body.Payment.CreditCard.Expiration,
	//SecurityCode:   body.Payment.CreditCard.CVV,
	//Brand:          body.Payment.CreditCard.Brand,
	//},
	//},
	//}

	//response, err := client.NewOrder(&order)

	//if err != nil {
	//c.AbortWithStatusJSON(400, gin.H{
	//"error": err,
	//})
	//return
	//}

	c.JSON(http.StatusOK, gin.H{
		"acquirer": "cielo",
		"order":    nil,
	})
}

// Capture an existing transaction
func (controller *AcquirerController) Capture(c *gin.Context) {
	response, err := controller.Cielo.CaptureOrder(c.Param("id"))

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
	response, err := controller.Cielo.GetOrder(c.Param("id"))

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
