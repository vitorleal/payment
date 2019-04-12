package stone

import (
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

	if s.OrderResult != nil {
		response.Id = s.OrderResult.OrderReference
		response.AuthorizationCode = s.OrderResult.OrderKey
	}

	// If CreditCard
	if len(s.CreditCardTransactionResultCollection) > 0 {
		transaction := s.CreditCardTransactionResultCollection[0]

		response.Amount = transaction.AmountInCents
		//response.CreditCard = &g.CreditCard{}
		response.NSU = transaction.UniqueSequentialNumber
		response.TID = transaction.TransactionIdentifier
	}

	// If BankingBillet
	if len(s.BoletoTransactionResultCollection) > 0 {
		transaction := s.BoletoTransactionResultCollection[0]

		response.Amount = transaction.AmountInCents
		response.BarCode = transaction.Barcode
		response.BoletoUrl = transaction.BoletoUrl
	}

	return response
}
