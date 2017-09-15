package errors

type RequestError struct {
	Message string `json:"message,omitempty"`
	Code    uint32 `json:"code,omitempty"`
}

var ()
