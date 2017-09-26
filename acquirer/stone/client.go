package stone

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/sling"
)

const Name = "stone"
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
func (client *Client) NewSale(sale *Sale) (*Sale, *Sale) {
	body, _ := json.Marshal(sale)
	fmt.Printf("%s", body)

	responseSale := new(Sale)
	responseError := new(Sale)

	_, _ = client.Api.Post(BasePath).BodyJSON(sale).Receive(responseSale, responseError)

	return responseSale, responseError
}

// Capture a stone sale
func (client *Client) CaptureSale(id string) (*Sale, error) {
	responseSale := new(Sale)
	_, err := client.Api.Get(BasePath + "/Capture").ReceiveSuccess(responseSale)

	return responseSale, err
}

// Get a stone sale
func (client *Client) GetSale(id string) (*Sale, *Sale) {
	responseSale := new(Sale)
	responseError := new(Sale)

	_, _ = client.Api.Get(BasePath+"/Query/OrderKey="+id).Receive(responseSale, responseError)

	return responseSale, responseError
}
