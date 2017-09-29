package stone

import (
	g "github.com/ingresse/payment/gateway"
)

type SaleResponse struct {
	Sale

	CreditCardTransactionResultCollection []*CreditCardTransaction `json:",omitempty"`
	BoletoTransactionResultCollection     []*BoletoTransaction     `json:",omitempty"`
	OrderResult                           *Order                   `json:",omitempty"`
	BuyerKey                              string                   `json:",omitempty"`
}

// FormatResponse use data from a SaleResponse to create a Gateway Payment
func (s *SaleResponse) FormatResponse() *g.Response {
	response := g.Response{}

	response.Acquirer = Name

	if s.OrderResult != nil {
		response.Id = s.OrderResult.OrderReference
		response.AuthorizationCode = s.OrderResult.OrderKey
	}

	if s.CreditCardTransactionResultCollection != nil {
		transaction := s.CreditCardTransactionResultCollection[0]

		response.Amount = transaction.AuthorizedAmountInCents
		response.CreditCard = &g.CreditCard{}
	}

	if s.BoletoTransactionResultCollection != nil {
		transaction := s.BoletoTransactionResultCollection[0]

		response.Amount = transaction.AuthorizedAmountInCents
	}

	return &response
}
