package cielo

type Sale struct {
	MerchantOrderId string    `json:",omitempty"`
	Customer        *Customer `json:",omitempty"`
	Payment         *Payment  `json:",omitempty"`
}

// -----------------------------

type Payment struct {
	Type              string            `json:",omitempty"`
	Amount            uint32            `josn:",omitempty"`
	Installments      uint32            `json:",omitempty"`
	SoftDescriptor    string            `json:",omitempty"`
	Currency          string            `json:",omitempty"`
	Country           string            `json:",omitempty"`
	Capture           bool              `json:",omitempty"`
	CreditCard        *CreditCard       `json:",omitempty"`
	Provider          string            `json:",omitempty"`
	Address           string            `json:",omitempty"`
	BoletoNumber      string            `json:",omitempty"`
	Assignor          string            `json:",omitempty"`
	Demonstrative     string            `json:",omitempty"`
	ExpirationDate    string            `json:",omitempty"`
	Identification    string            `json:",omitempty"`
	Instructions      string            `json:",omitempty"`
	ProofOfSale       string            `json:",omitempty"`
	ServiceTaxAmount  uint32            `json:",omitempty"`
	Tid               string            `json:",omitempty"`
	Authorizationcode string            `json:",omitempty"`
	PaymentId         string            `json:",omitempty"`
	ECI               string            `json:",omitempty"`
	CapturedAmount    uint32            `json:",omitempty"`
	Status            uint32            `json:",omitempty"`
	ReturnCode        string            `json:",omitempty"`
	ReturnMessage     string            `json:",omitempty"`
	Links             []Link            `json:",omitempty"`
	RecurrentPayment  *RecurrentPayment `json:",omitempty"`
}

// -----------------------------

type CreditCard struct {
	CardNumber     string
	Holder         string
	ExpirationDate string
	SecurityCode   string
	Brand          string
	CardToken      string `json:",omitempty"`
	SaveCard       bool   `json:",omitempty"`
}

type Link struct {
	Method string `json:",omitempty"`
	Rel    string `json:",omitempty"`
	Href   string `json:",omitempty"`
}

type RecurrentPayment struct {
	AuthorizeNow bool   `json:",omitempty"`
	EndDate      string `json:",omitempty"`
	Interval     string `json:",omitempty"`
}

// -----------------------------

type Customer struct {
	Name            string
	Email           string   `json:",omitempty"`
	Birthdate       string   `json:",omitempty"`
	Address         *Address `json:",omitempty"`
	DeliveryAddress *Address `json:",omitempty"`
}

type Address struct {
	Street     string `json:",omitempty"`
	Number     string `json:",omitempty"`
	Complement string `json:",omitempty"`
	ZipCode    string `json:",omitempty"`
	City       string `json:",omitempty"`
	State      string `json:",omitempty"`
	Country    string `json:",omitempty"`
}
