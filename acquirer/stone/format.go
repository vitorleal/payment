package stone

import (
	"github.com/ingresse/payment/gateway"
)

// Create a new sale from a gateway payment
func FormatSale(payment *gateway.Payment) *Sale {
	// Create a new sale
	newSale := Sale{
		Order: &Order{
			OrderReference: payment.Id,
		},
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
		DocumentType:      "CPF",
		DocumentNumber:    payment.Customer.Document,
		AddressCollection: []*Address{address},
	}

	newSale.Buyer = buyer

	// If payment with CreditCard
	if payment.WithCredicard() {
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
		cardTransaction.CreditCardOperation = "AuthOnly"

		// Stone homolog variable
		cardTransaction.Options = &Options{
			PaymentMethodCode: 1,
		}

		newSale.CreditCardTransactionCollection = []*CreditCardTransaction{&cardTransaction}
	}

	// If payment with BankingBillet
	if payment.WithBankBillet() {
		bankBilling := BoletoTransaction{
			AmountInCents:        payment.Amount,
			BankNumber:           "033",
			Instructions:         payment.BankingBillet.Instructions,
			TransactionReference: "BoletoComBuyerRemessa",
		}

		// BankingBillet options
		bankBilling.Options = &Options{
			PaymentMethodCode:               1,
			DaysToAddInBoletoExpirationDate: 5,
		}
	}

	return &newSale
}
