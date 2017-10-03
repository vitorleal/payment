package stone

import (
	g "github.com/ingresse/payment/gateway"
)

type (
	// Sale represents the stone sale data
	Sale struct {
		MerchantKey  string `json:",omitempty"`
		RequestKey   string `json:",omitempty"`
		InternalTime uint32 `json:",omitempty"`

		CreditCardTransactionCollection []*CreditCardTransaction `json:",omitempty"`
		BoletoTransactionCollection     []*BoletoTransaction     `json:",omitempty"`
		Order                           *Order                   `json:",omitempty"`
		Buyer                           *Buyer                   `json:",omitempty"`
	}

	// CreditCardTransaction represents the CreditCard transaction data for a sale
	CreditCardTransaction struct {
		AcquirerMessage             string `json:",omitempty"`
		AcquirerName                string `json:",omitempty"`
		AcquirerReturnCode          string `json:",omitempty"`
		AffiliationCode             string `json:",omitempty"`
		AuthorizationCode           string `json:",omitempty"`
		CapturedDate                string `json:",omitempty"`
		Success                     bool   `json:",omitempty"`
		TransactionIdentifier       string `json:",omitempty"`
		TransactionKey              string `json:",omitempty"`
		TransactionKeyToAcquirer    string `json:",omitempty"`
		TransactionReference        string `json:",omitempty"`
		UniqueSequentialNumber      string `json:",omitempty"`
		CreditCardOperation         string `json:",omitempty"`
		CreditCardTransactionStatus string `json:",omitempty"`
		DueDate                     string `json:",omitempty"`
		ThirdPartyAffiliationCode   string `json:",omitempty"`
		RefundedAmountInCents       uint32 `json:",omitempty"`
		PaymentMethodName           string `json:",omitempty"`
		VoidedAmountInCents         uint32 `json:",omitempty"`

		AmountInCents           uint32      `json:",omitempty"`
		AuthorizedAmountInCents uint32      `json:",omitempty"`
		CreditCard              *CreditCard `json:",omitempty"`
		InstallmentCount        uint32      `json:",omitempty"`
		Options                 *Options    `json:",omitempty"`
	}

	// CreditCard represent the creditcard data for a sale
	CreditCard struct {
		InstantBuyKey          string `json:",omitempty"`
		IsExpiredCreditCard    bool   `json:",omitempty"`
		MaskedCreditCardNumber string `json:",omitempty"`

		CreditCardBrand  string `json:",omitempty"`
		CreditCardNumber string `json:",omitempty"`
		ExpMonth         string `json:",omitempty"`
		ExpYear          string `json:",omitempty"`
		HolderName       string `json:",omitempty"`
		SecurityCode     string `json:",omitempty"`
	}

	// BoletoTransaction represents the boleto transaction for a sale
	BoletoTransaction struct {
		AcquirerReturnCode      string `json:",omitempty"`
		AcquirerReturnMessage   string `json:",omitempty"`
		Barcode                 string `json:",omitempty"`
		BoletoTransactionStatus string `json:",omitempty"`
		BoletoUrl               string `json:",omitempty"`
		DocumentNumber          string `json:",omitempty"`
		NossoNumero             string `json:",omitempty"`
		Success                 bool   `json:",omitempty"`
		TransactionKey          string `json:",omitempty"`
		TransactionReference    string `json:",omitempty"`

		AmountInCents           uint32   `json:",omitempty"`
		AuthorizedAmountInCents uint32   `json:",omitempty"`
		BankNumber              string   `json:",omitempty"`
		Instructions            string   `json:",omitempty"`
		Options                 *Options `json:",omitempty"`
	}

	// Options represents payment options for a sale
	Options struct {
		DaysToAddInBoletoExpirationDate uint8 `json:",omitempty"`
		PaymentMethodCode               uint8 `json:",omitempty"`
	}

	// Order represents the order data for a sale
	Order struct {
		OrderReference string `json:",omitempty"`
		CreateDate     string `json:",omitempty"`
		OrderKey       string `json:",omitempty"`
	}

	// Buyer represents the data of the buyer
	Buyer struct {
		AddressCollection []*Address `json:",omitempty"`
		DocumentNumber    string     `json:",omitempty"`
		DocumentType      string     `json:",omitempty"`
		Name              string     `json:",omitempty"`
		PersonType        string     `json:",omitempty"`
	}

	// Address represnets the address data of a buyer
	Address struct {
		AddressType string `json:",omitempty"`
		City        string `json:",omitempty"`
		Complement  string `json:",omitempty"`
		Country     string `json:",omitempty"`
		District    string `json:",omitempty"`
		Number      string `json:",omitempty"`
		State       string `json:",omitempty"`
		Street      string `json:",omitempty"`
		ZipCode     string `json:",omitempty"`
	}
)

// FormatPayment use a gateway payment data for fill up a Stone sale
func (s *Sale) FromPayment(payment *g.Payment) {
	s.Order = &Order{
		OrderReference: payment.Id,
	}

	// Create the customer address
	address := &Address{}

	if payment.Customer.Address != nil {
		address.AddressType = "Residential"
		address.City = payment.Customer.Address.City
		address.Complement = payment.Customer.Address.Complement
		address.Country = payment.Customer.Address.Country
		address.Number = payment.Customer.Address.Number
		address.State = payment.Customer.Address.State
		address.Street = payment.Customer.Address.Street
		address.ZipCode = payment.Customer.Address.ZipCode
	}

	// Create the customer
	buyer := &Buyer{
		Name:              payment.Customer.Name,
		DocumentType:      CPF,
		DocumentNumber:    payment.Customer.Document,
		AddressCollection: []*Address{address},
	}

	s.Buyer = buyer

	// If payment with CreditCard
	if payment.WithCrediCard() {
		cardTransaction := CreditCardTransaction{
			AmountInCents: payment.Amount,
			CreditCard: &CreditCard{
				CreditCardBrand:  payment.CreditCard.Brand,
				CreditCardNumber: payment.CreditCard.Number,
				HolderName:       payment.CreditCard.Holder,
				ExpMonth:         payment.CreditCard.Expiration[:2],
				ExpYear:          payment.CreditCard.Expiration[3:],
				SecurityCode:     payment.CreditCard.CVV,
			},
		}

		// Do not capture the transaction, only authorize
		cardTransaction.CreditCardOperation = Authorize

		// Stone homolog variable
		cardTransaction.Options = &Options{
			PaymentMethodCode: 1,
		}

		s.CreditCardTransactionCollection = []*CreditCardTransaction{&cardTransaction}
	}

	// If payment with BankingBillet
	if payment.WithBankingBillet() {
		bankingBillet := BoletoTransaction{
			AmountInCents:        payment.Amount,
			Instructions:         payment.BankingBillet.Instructions,
			TransactionReference: payment.Id,
		}

		// BankingBillet options
		bankingBillet.Options = &Options{
			PaymentMethodCode:               1,
			DaysToAddInBoletoExpirationDate: payment.BankingBillet.Expiration,
		}

		s.BoletoTransactionCollection = []*BoletoTransaction{&bankingBillet}
	}
}
