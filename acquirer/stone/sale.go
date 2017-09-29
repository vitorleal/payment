package stone

import (
	g "github.com/ingresse/payment/gateway"
)

type Sale struct {
	MerchantKey  string `json:",omitempty"`
	RequestKey   string `json:",omitempty"`
	InternalTime uint32 `json:",omitempty"`

	CreditCardTransactionCollection []*CreditCardTransaction `json:",omitempty"`
	BoletoTransactionCollection     []*BoletoTransaction     `json:",omitempty"`
	Order                           *Order                   `json:",omitempty"`
	Buyer                           *Buyer                   `json:",omitempty"`
}

type CreditCardTransaction struct {
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

type CreditCard struct {
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

type BoletoTransaction struct {
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

type Options struct {
	DaysToAddInBoletoExpirationDate uint8 `json:",omitempty"`
	PaymentMethodCode               uint8 `json:",omitempty"`
}

type Order struct {
	OrderReference string `json:",omitempty"`
	CreateDate     string `json:",omitempty"`
	OrderKey       string `json:",omitempty"`
}

type Buyer struct {
	AddressCollection []*Address `json:",omitempty"`
	DocumentNumber    string     `json:",omitempty"`
	DocumentType      string     `json:",omitempty"`
	Name              string     `json:",omitempty"`
	PersonType        string     `json:",omitempty"`
}

type Address struct {
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

// FormatPayment use a Gateway payment data for fill up a Stone sale
func (s *Sale) FromPayment(payment *g.Payment) {
	s.Order = &Order{
		OrderReference: payment.Id,
	}

	// Create the customer address
	address := &Address{
		AddressType: "Residential",
		City:        payment.Customer.Address.City,
		Complement:  payment.Customer.Address.Complement,
		Country:     payment.Customer.Address.Country,
		Number:      payment.Customer.Address.Number,
		State:       payment.Customer.Address.State,
		Street:      payment.Customer.Address.Street,
		ZipCode:     payment.Customer.Address.ZipCode,
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
	if payment.WithBankBillet() {
		bankBilling := BoletoTransaction{
			AmountInCents:        payment.Amount,
			Instructions:         payment.BankingBillet.Instructions,
			TransactionReference: payment.Id,
		}

		// BankingBillet options
		bankBilling.Options = &Options{
			PaymentMethodCode:               1,
			DaysToAddInBoletoExpirationDate: payment.BankingBillet.Expiration,
		}
	}
}
