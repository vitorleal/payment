package cielo

// Order base
type Order struct {
	MerchantOrderId string    `json:",omitempty"`
	Customer        *Customer `json:",omitempty"`
	Payment         *Payment  `json:",omitempty"`
}
