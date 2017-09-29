package acquirer

import (
	"github.com/ingresse/payment/acquirer/stone"
	"github.com/ingresse/payment/gateway"
)

// Acquirer Client interface
type ClientInterface interface {
	Get(id string) (*gateway.Response, error)

	Authorize(payment *gateway.Payment) (*gateway.Response, error)
	Capture(id string) (*gateway.Response, error)
	Cancel(payment *gateway.Payment) (*gateway.Response, error)
}

type Acquirer struct {
	Client ClientInterface
}

// Create new Acquirer based on the acquirer name
func NewAcquirer(name string) *Acquirer {
	return &Acquirer{
		Client: stone.NewClient(stone.Merchant{
			Key: "f2a1f485-cfd4-49f5-8862-0ebc438ae923",
		}, stone.Production),
	}
}

// Authorize a payment using the Acquirer client
func (acquirer *Acquirer) Authorize(payment *gateway.Payment) (*gateway.Response, error) {
	return acquirer.Client.Authorize(payment)
}

// Capture an autorized payment using the Acquirer client
func (acquirer *Acquirer) Capture(id string) (*gateway.Response, error) {
	return acquirer.Client.Capture(id)
}

// Get a payment using the Acquirer client
func (acquirer *Acquirer) Get(id string) (*gateway.Response, error) {
	return acquirer.Client.Get(id)
}
