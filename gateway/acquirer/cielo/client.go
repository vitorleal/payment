package cielo

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/satori/go.uuid"
)

const Name = "cielo"
const BasePath = "/1/sales/"

var (
	Production = Environment{
		Url:   "https://api.cieloecommerce.cielo.com.br",
		Query: "https://apiquery.cieloecommerce.cielo.com.br",
	}

	Sandbox = Environment{
		Url:   "https://apisandbox.cieloecommerce.cielo.com.br",
		Query: "https://apiquerysandbox.cieloecommerce.cielo.com.br",
	}
)

// Cielo merchant
type Merchant struct {
	Id, Key string
}

// Cielo environment
type Environment struct {
	Url, Query string
}

// Cielo client
type Client struct {
	Api      *sling.Sling
	Query    *sling.Sling
	Merchant Merchant
	Env      Environment
}

// New Cielo API clients
func New(merchant Merchant, env Environment) *Client {
	api := sling.New().Client(nil)

	api.Add("Accept", "application/json")
	api.Add("Accept-Encoding", "gzip")
	api.Add("Content-Type", "application/json")
	api.Add("User-Agent", "Payment/1.0")
	api.Add("MerchantId", merchant.Id)
	api.Add("MerchantKey", merchant.Key)
	api.Add("RequestId", uuid.NewV4().String())

	client := Client{
		Api:      api.New().Base(env.Url),
		Query:    api.New().Base(env.Query),
		Merchant: merchant,
		Env:      env,
	}

	return &client
}

// Create a cielo sale
func (client *Client) SaleNew(sale *Sale) (*Sale, error) {
	body, err := json.Marshal(sale)
	fmt.Printf("%s", body)

	responseSale := new(Sale)
	_, err = client.Api.Post(BasePath).BodyJSON(sale).ReceiveSuccess(responseSale)

	return responseSale, err
}

// Capture a cielo sale
func (client *Client) SaleCapture(id string) (*Sale, error) {
	responseSale := new(Sale)
	_, err := client.Api.Put(BasePath + id + "/capture").ReceiveSuccess(responseSale)

	return responseSale, err
}

// Get a cielo sale
func (client *Client) SaleGet(id string) (*Sale, error) {
	responseSale := new(Sale)
	_, err := client.Query.Get(BasePath + id).ReceiveSuccess(responseSale)

	return responseSale, err
}
