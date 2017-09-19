package stone

type Sale struct {
	MerchantKey  string `json:",omitempty"`
	RequestKey   string `json:",omitempty"`
	InternalTime uint32 `json:",omitempty"`

	ErrorReport interface{} `json:",omitempty"`

	CreditCardTransactionCollection []*CreditCardTransaction `json:",omitempty"`
	BoletoTransactionCollection     []*BoletoTransaction     `json:",omitempty"`
	Order                           *Order                   `json:",omitempty"`
	Buyer                           *Buyer                   `json:",omitempty"`
}

// -----------------------------

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

	AmountInCents    uint32      `json:",omitempty"`
	CreditCard       *CreditCard `json:",omitempty"`
	InstallmentCount uint32      `json:",omitempty"`
	Options          *Options    `json:",omitempty"`
}

type CreditCard struct {
	InstantBuyKey          string `json:",omitempty"`
	IsExpiredCreditCard    string `json:",omitempty"`
	MaskedCreditCardNumber string `json:",omitempty"`

	CreditCardBrand  string `json:",omitempty"`
	CreditCardNumber string `json:",omitempty"`
	ExpMonth         string `json:",omitempty"`
	ExpYear          string `json:",omitempty"`
	HolderName       string `json:",omitempty"`
	SecurityCode     string `json:",omitempty"`
}

// -----------------------------

type BoletoTransaction struct {
	AmountInCents        uint32   `json:",omitempty"`
	BankNumber           string   `json:",omitempty"`
	Instructions         string   `json:",omitempty"`
	TransactionReference string   `json:",omitempty"`
	Options              *Options `json:",omitempty"`
}

// -----------------------------

type Options struct {
	DaysToAddInBoletoExpirationDate uint32 `json:",omitempty"`
	PaymentMethodCode               uint32 `json:",omitempty"`
}

// -----------------------------

type Order struct {
	OrderReference string `json:",omitempty"`
	CreateDate     string `json:",omitempty"`
	OrderKey       string `json:",omitempty"`
}

// -----------------------------

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
