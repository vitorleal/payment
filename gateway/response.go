package gateway

type Response struct {
	Payment
	NSU               string `json:"nsu"`
	AuthorizationCode string `json:"authorizationCode"`
	Token             string `json:"token"`
}
