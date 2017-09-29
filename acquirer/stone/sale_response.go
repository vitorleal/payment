package stone

import (
	"fmt"
	g "github.com/ingresse/payment/gateway"
)

// SaleResponse represents an Stone api response
type SaleResponse struct {
	Sale

	CreditCardTransactionResultCollection []*CreditCardTransaction `json:",omitempty"`
	BoletoTransactionResultCollection     []*BoletoTransaction     `json:",omitempty"`
	OrderResult                           *Order                   `json:",omitempty"`
	BuyerKey                              string                   `json:",omitempty"`
}

// FormatResponse use data from a SaleResponse to create a Gateway Payment
func (s *SaleResponse) FormatResponse() *g.Response {
	response := new(g.Response)
	response.Acquirer = Name

	fmt.Printf("%+v", s)

	if s.OrderResult != nil {
		response.Id = s.OrderResult.OrderReference
		response.AuthorizationCode = s.OrderResult.OrderKey
	}

	// If CreditCard
	if len(s.CreditCardTransactionResultCollection) > 0 {
		transaction := s.CreditCardTransactionResultCollection[0]

		response.Amount = transaction.AuthorizedAmountInCents
		response.CreditCard = &g.CreditCard{}
	}

	// If BankingBillet
	if len(s.BoletoTransactionResultCollection) > 0 {
		transaction := s.BoletoTransactionResultCollection[0]

		response.Amount = transaction.AuthorizedAmountInCents
	}

	return response
}
