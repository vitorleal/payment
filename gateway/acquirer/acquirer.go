package acquirer

import (
	g "github.com/ingresse/payment/gateway"
	"github.com/ingresse/payment/gateway/acquirer/stone"
)

// ClientInterface is the interface for differents acquirers
type ClientInterface interface {
	Get(id string) (*g.Response, *g.Error)
	Capture(id string) (*g.Response, *g.Error)
	Authorize(payment *g.Payment) (*g.Response, *g.Error)
	Cancel(payment *g.Payment) (*g.Response, *g.Error)
}

// Acquirer represents the struct of the acquirers
type Acquirer struct {
	Client ClientInterface
}

// NewAcquirer create a new Acquirer based on the acquirer name
func NewAcquirer(name string) *Acquirer {
	return &Acquirer{
		Client: stone.NewClient(stone.Merchant{
			Key: "f2a1f485-cfd4-49f5-8862-0ebc438ae923",
		}, stone.Production),
	}
}

// Authorize will autorize a payment using the Acquirer client
func (a *Acquirer) Authorize(payment *g.Payment) (*g.Response, *g.Error) {
	return a.Client.Authorize(payment)
}

// Capture will capture an autorized payment using the Acquirer client
func (a *Acquirer) Capture(id string) (*g.Response, *g.Error) {
	return a.Client.Capture(id)
}

// Get will a payment using the Acquirer client
func (a *Acquirer) Get(id string) (*g.Response, *g.Error) {
	return a.Client.Get(id)
}

// Cancel will cancel an authorized or payed payment using the Acquirer client
func (a *Acquirer) Cancel(payment *g.Payment) (*g.Response, *g.Error) {
	return a.Client.Cancel(payment)
}
