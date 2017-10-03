package gateway

type (
	// Payment represnet a new request payment data
	Payment struct {
		Id            string         `json:"id,omitempty"            binding:"required"`
		Acquirer      string         `json:"acquirer,omitempty"      binding:"required,acquirer"`
		Antifraud     []*Antifraud   `json:"antifraud,omitempty"     binding:"dive"`
		MerchantId    string         `json:"merchantId,omitempty"    binding:"omitempty"`
		Amount        uint32         `json:"amount,omitempty"        binding:"exists,ne=0"`
		Interests     uint32         `json:"interests,omitempty"`
		Customer      *Customer      `json:"customer,omitempty"      binding:"required"`
		Items         []*Item        `json:"items,omitempty"         binding:"dive"`
		CreditCard    *CreditCard    `json:"creditCard,omitempty"    binding:"omitempty"`
		BankingBillet *BankingBillet `json:"bankingBillet,omitempty" binding:"omitempty"`
	}

	// CreditCard represents the creditCard data for the payment
	CreditCard struct {
		SoftDescriptor string `json:"softDescriptor,omitempty"`
		Installments   uint8  `json:"installments,omitempty" binding:"required,min=1"`
		Number         string `json:"number,omitempty"       binding:"required,min=10,max=24"`
		Holder         string `json:"holder,omitempty"       binding:"required"`
		Expiration     string `json:"expiration,omitempty"   binding:"required"`
		CVV            string `json:"cvv,omitempty"          binding:"required"`
		Brand          string `json:"brand,omitempty"        binding:"required,creditCardBrand"`
		SaveCard       bool   `json:"saveCard,omitempty"     binding:"omitempty"`

		Token  string `json:"token,omitempty"  binding:"omitempty"`
		Masked string `json:"masked,omitempty" binding:"omitempty"`
	}

	// BankingBillet represents the bankingBillet data for the payment
	BankingBillet struct {
		Expiration   uint8  `json:"expiration"   binding:"required,min=3,max=6"`
		Instructions string `json:"instructions" binding:"required"`
	}

	// Antifraud represents the antifraud data for the payment
	Antifraud struct {
		Name  string  `json:"name,omitempty", binding:"required,antifraud"`
		Score float64 `json:"score,omitempty" binding:"required,ne=0"`
	}

	// Customer represents the customer data for the payment
	Customer struct {
		Id       string `json:"id"       binding:"required"`
		Name     string `json:"name"     binding:"required"`
		Document string `json:"document" binding:"required"`

		Address *Address `json:"address,omitempty"`
	}

	// Address represents the customer address data
	Address struct {
		Street     string `json:"street,omitempty"     binding:"required"`
		Number     string `json:"number,omitempty"     binding:"required"`
		Complement string `json:"complement,omitempty" binding:"required"`
		ZipCode    string `json:"zipcode,omitempty"    binding:"required"`
		City       string `json:"city,omitempty"       binding:"required"`
		State      string `json:"state,omitempty"      binding:"required"`
		Country    string `json:"country,omitempty"    binding:"required"`
	}

	// Item represents the products related to the payment
	Item struct {
		Id        string `json:"id,omitempty"        binding:"required"`
		Type      string `json:"type,omitempty"      binding:"required"`
		Name      string `json:"name,omitempty"      binding:"required"`
		UnitPrice uint32 `json:"unitPrice,omitempty" binding:"required"`
		Quantity  uint32 `json:"quantity,omitempty"  binding:"required"`
	}
)

// WithCrediCard verify if payment is with creditcard
func (payment *Payment) WithCrediCard() bool {
	if payment.CreditCard != nil {
		return true
	}

	return false
}

// WithBankingBillet verify if payment with bankingBillet
func (payment *Payment) WithBankingBillet() bool {
	if payment.BankingBillet != nil && !payment.WithCrediCard() {
		return true
	}

	return false
}
