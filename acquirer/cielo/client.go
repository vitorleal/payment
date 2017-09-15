package cielo

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/satori/go.uuid"
)

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

var basePath = "/1/sales/"

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

// New Cielo APIs clients
func New(merchant Merchant, env Environment) *Client {
	api := sling.New().Client(nil)

	api.Add("Accept", "application/json")
	api.Add("Accept-Encoding", "gzip")
	api.Add("Content-Type", "application/json")
	api.Add("User-Agent", "Ingresse-Payment/1.0")
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

// Create a cielo order
func (client *Client) NewOrder(order *Order) (*Order, error) {
	body, err := json.Marshal(order)
	fmt.Printf("%s", body)

	responseOrder := new(Order)
	_, err = client.Api.Post(basePath).BodyJSON(order).ReceiveSuccess(responseOrder)

	return responseOrder, err
}

// Capture a cielo order
func (client *Client) CaptureOrder(id string) (*Order, error) {
	responseOrder := new(Order)
	_, err := client.Api.Put(basePath + id + "/capture").ReceiveSuccess(responseOrder)

	return responseOrder, err
}

// Get a cielo order
func (client *Client) GetOrder(id string) (*Order, error) {
	responseOrder := new(Order)
	_, err := client.Query.Get(basePath + id).ReceiveSuccess(responseOrder)

	return responseOrder, err
}
