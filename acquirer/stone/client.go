package stone

import (
	"github.com/dghubble/sling"
	e "github.com/ingresse/payment/errors"
	g "github.com/ingresse/payment/gateway"
)

const BasePath = "/Sale"

var Production = Environment{
	Url: "https://transaction.stone.com.br",
}

type (
	// Merchant represents the stone merchant information
	Merchant struct {
		Key string
	}

	// Environment represent the stone environment
	Environment struct {
		Url string
	}

	// Clinet implements Acquirer ClientInterface
	Client struct {
		Api      *sling.Sling
		Merchant Merchant
		Env      Environment
	}
)

// New client create a stone client
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

// Authorize will authorize a sale in Stone
func (client *Client) Authorize(payment *g.Payment) (*g.Response, *e.ApiError) {
	body := new(Sale)
	body.FromPayment(payment)

	sale := new(SaleResponse)
	saleError := new(SaleError)

	client.Api.Post(BasePath).BodyJSON(body).Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(AuthorizeError, saleError.ErrorReport, AuthorizeErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}

// Capture will capture an authorized sale in Stone
func (client *Client) Capture(id string) (*g.Response, *e.ApiError) {
	sale := new(SaleResponse)
	saleError := new(SaleError)

	client.Api.Get(BasePath+"/Capture").Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(CaptureError, saleError.ErrorReport, CaptureErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}

// Get will get sale information in Stone
func (client *Client) Get(id string) (*g.Response, *e.ApiError) {
	sale := new(SaleDataResponse)
	saleError := new(SaleError)

	client.Api.Get(BasePath+"/Query/OrderKey="+id).Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(GetSaleError, saleError.ErrorReport, GetSaleErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}

// Cancel will cancel an authorized or payed sale in Stone
func (client *Client) Cancel(payment *g.Payment) (*g.Response, *e.ApiError) {
	body := new(Sale)
	body.FromPayment(payment)

	sale := new(SaleResponse)
	saleError := new(SaleError)

	client.Api.Post(BasePath+"/Cancel").BodyJSON(body).Receive(sale, saleError)

	// If error
	if saleError.ErrorReport != nil {
		err := BadRequestError(CancelError, saleError.ErrorReport, CancelErrorCode)
		return nil, err
	}

	response := sale.FormatResponse()
	return response, nil
}
