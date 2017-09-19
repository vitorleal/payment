package acquirer

import (
	"github.com/ingresse/payment/gateway"
)

// Acquirer interface
type Acquirer interface {
	New() *Acquirer
	NewSale(payment *gateway.Payment) (*AcquirerResponse, error)
	CaptureSale(id string) (*AcquirerResponse, error)
	GetSale(id string) (*AcquirerResponse, error)
}

type AcquirerResponse struct{}
