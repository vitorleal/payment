package gateway

// Response represent the response data for a payment
type Response struct {
	Payment

	// CreditCard response
	NSU               string `json:"nsu,omitempty"`
	TID               string `json:"tid,omitempty"`
	AuthorizationCode string `json:"authorizationCode,omitempty"`
	Token             string `json:"token,omitempty"`

	// BoletoResponse
	BarCode      string `json:"barCode,omitempty"`
	BoletoStatus string `json:"boletoStatus,omitempty"`
	BoletoUrl    string `json:"boletoUrl,omitempty"`
}
