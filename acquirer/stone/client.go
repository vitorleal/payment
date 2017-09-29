package stone

import (
	"github.com/dghubble/sling"
	g "github.com/ingresse/payment/gateway"
)

const BasePath = "/Sale"

var Production = Environment{
	Url: "https://transaction.stone.com.br",
}

// Stone merchant
type Merchant struct {
	Key string
}

// Stone environment
type Environment struct {
	Url string
}

// Stone client
// implements Acquirer ClientInterface
type Client struct {
	Api      *sling.Sling
	Merchant Merchant
	Env      Environment
}

// Create a stone client
func NewClient(merchant Merchant, env Environment) *Client {
	api := sling.New().Client(nil)

	api.Add("Accept", "application/json")
	api.Add("Accept-Encoding", "gzip")
	api.Add("Content-Type", "application/json")
	api.Add("User-Agent", "Ingresse-Payment/1.0")
	api.Add("MerchantKey", merchant.Key)

	client := Client{
		Api:      api.New().Base(env.Url),
		Merchant: merchant,
		Env:      env,
	}

	return &client
}

// Authorize authorize a sale in Stone
func (client *Client) Authorize(payment *g.Payment) (*g.Response, error) {
	body := Sale{}
	body.FromPayment(payment)

	sale := SaleResponse{}
	saleError := SaleError{}

	client.Api.Post(BasePath).BodyJSON(body).Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(AuthorizeError, saleError.ErrorReport, AuthorizeErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}

// Capture capture an authorized sale in Stone
func (client *Client) Capture(id string) (*g.Response, error) {
	sale := SaleResponse{}
	saleError := SaleError{}

	client.Api.Get(BasePath+"/Capture").Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(CaptureError, saleError.ErrorReport, CaptureErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}

// Get get sale information in Stone
func (client *Client) Get(id string) (*g.Response, error) {
	sale := SaleDataResponse{}
	saleError := SaleError{}

	client.Api.Get(BasePath+"/Query/OrderKey="+id).Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(GetSaleError, saleError.ErrorReport, GetSaleErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}

// Cancel cancel an authorized or payed sale in Stone
func (client *Client) Cancel(payment *g.Payment) (*g.Response, error) {
	body := Sale{}
	body.FromPayment(payment)

	sale := SaleResponse{}
	saleError := SaleError{}

	client.Api.Post(BasePath+"/Cancel").BodyJSON(body).Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(CancelError, saleError.ErrorReport, CancelErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}
