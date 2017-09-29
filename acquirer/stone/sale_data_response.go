package stone

import (
	g "github.com/ingresse/payment/gateway"
)

type SaleDataResponse struct {
	Sale

	SaleDataCollection []*SaleData `json:",omitempty"`
	SaleDataCount      uint32      `json:",omitempty"`
}

type SaleData struct {
	CreditCardTransactionDataCollection []*CreditCardTransaction `json:",omitempty"`
	BoletoTransactionDataCollection     []*BoletoTransaction     `json:",omitempty"`
	OrderData                           *Order                   `json:",omitempty"`
	BuyerKey                            string                   `json:",omitempty"`
}

// FormatResponse use data from a SaleDataResponse to create a Gateway Payment
func (s *SaleDataResponse) FormatResponse() *g.Response {
	response := g.Response{}

	return &response
}
