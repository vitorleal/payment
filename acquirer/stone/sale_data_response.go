package stone

import (
	g "github.com/ingresse/payment/gateway"
)

type (
	// SaleDataResponse represents the sale response when
	// getting the information in the stone api
	SaleDataResponse struct {
		Sale

		SaleDataCollection []*SaleData `json:",omitempty"`
		SaleDataCount      uint32      `json:",omitempty"`
	}

	// SaleData represnet the sale data information when
	// getting the information in the stone api
	SaleData struct {
		CreditCardTransactionDataCollection []*CreditCardTransaction `json:",omitempty"`
		BoletoTransactionDataCollection     []*BoletoTransaction     `json:",omitempty"`
		OrderData                           *Order                   `json:",omitempty"`
		BuyerKey                            string                   `json:",omitempty"`
	}
)

// FormatResponse use data from a SaleDataResponse to create a Gateway Payment
func (s *SaleDataResponse) FormatResponse() *g.Response {
	response := new(g.Response)

	return response
}
