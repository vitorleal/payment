package gateway

import (
	"github.com/ingresse/payment/antifraud"
	"strings"
)

const (
	Cielo = "cielo"
	Stone = "stone"
)

type Payment struct {
	Id            string              `json:"id" binding:"required"`
	Acquirer      string              `json:"acquirer" binding:"required"`
	Antifraud     []antifraud.Service `json:"antifraud,omitempty" binding:"omitempty"`
	MerchantId    string              `json:"merchantId,omitempty" binding:"omitempty"`
	Type          string              `json:"type" binding:"required"`
	Amount        uint32              `json:"amount" binding:"exists,ne=0"`
	Interests     uint32              `json:"interests,omitempty"`
	Customer      *Customer           `json:"customer" binding:"required"`
	Items         []Item              `json:"items,omitempty" binding="omitempty"`
	CreditCard    *CreditCard         `json:"creditCard,omitempty" binding:"omitempty"`
	BankingBillet *BankingBillet      `json:"bankingBillet,omitempty" binding:"omitempty"`
}

type CreditCard struct {
	SoftDescriptor string `json:"softDescriptor" binding:"required"`
	Installments   uint32 `json:"installments" binding:"required"`
	Number         string `json:"number" binding:"required"`
	Holder         string `json:"holder" binding:"required"`
	Expiration     string `json:"expiration" binding:"required"`
	CVV            string `json:"cvv" binding:"required"`
	Brand          string `json:"brand" binding:"required"`
	SaveCard       bool   `json:"saveCard" binding:"omitempty"`
}

type BankingBillet struct {
	Provider     string `json:"provider" binding:"required"`
	Assignor     string `json:"assignor" binding:"required"`
	Expiration   string `json:"expiration" binding:"required"`
	Instructions string `json:"instructions" binding:"required"`
}

type Customer struct {
	Id       string   `json:"id" binding:"required"`
	Name     string   `json:"name" binding:"required"`
	Document string   `json:"document" binding:"required"`
	Address  *Address `json:"address,omitempty"`
}

type Address struct {
	Street     string `json:"street,omitempty"`
	Number     string `json:"number,omitempty"`
	Complement string `json:"complement,omitempty"`
	ZipCode    string `json:"zipcode,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
}

type Item struct {
	Id        string `json:"id,omitempty"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	UnitPrice uint32 `json:"unitPrice,omitempty"`
	Quantity  uint32 `json:"quantity,omitempty"`
}

// Check if payment accquirer is Stone
func (payment *Payment) IsStone() bool {
	if strings.ToLower(payment.Acquirer) == strings.ToLower(Stone) {
		return true
	}

	return false
}

// Check if payment acquirer is Cielo
func (payment *Payment) IsCielo() bool {
	if strings.ToLower(payment.Acquirer) == strings.ToLower(Cielo) {
		return true
	}

	return false
}
