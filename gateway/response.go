package gateway

// Response represent the response data for a payment
type Response struct {
	Payment

	NSU               string `json:"nsu"`
	AuthorizationCode string `json:"authorizationCode"`
	Token             string `json:"token"`
}
