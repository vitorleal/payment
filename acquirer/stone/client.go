package stone

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/sling"
	//"github.com/ingresse/payment/gateway"
)

var Production = Environment{
	Url: "https://transaction.stone.com.br",
}

const basePath = "/Sale"

// Stone merchant
type Merchant struct {
	Key string
}

// Stone environment
type Environment struct {
	Url string
}

// Stone client
type Client struct {
	Api      *sling.Sling
	Merchant Merchant
	Env      Environment
}

// Create a stone sale
func New(merchant Merchant, env Environment) *Client {
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

// Create a stone sale
func (client *Client) NewSale(sale *Sale) (*Sale, error) {
	body, err := json.Marshal(sale)
	fmt.Printf("%s", body)

	responseSale := new(Sale)
	_, err = client.Api.Post(basePath).BodyJSON(sale).ReceiveSuccess(responseSale)

	fmt.Printf("%+v", responseSale)

	return responseSale, err
}

// Capture a stone sale
func (client *Client) CaptureSale(id string) (*Sale, error) {
	responseSale := new(Sale)
	_, err := client.Api.Get(basePath + "/Capture").ReceiveSuccess(responseSale)

	return responseSale, err
}

// Get a stone sale
func (client *Client) GetSale(id string) (*Sale, error) {
	responseSale := new(Sale)

	_, err := client.Api.Get(basePath + "/Query/OrderKey=" + id).ReceiveSuccess(responseSale)

	return responseSale, err
}
