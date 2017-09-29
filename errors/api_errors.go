package errors

type ApiError struct {
	Status  int     `json:"status,omitempty"`
	Message string  `json:"message,omitempty"`
	Code    int     `json:"code,omitempty"`
	Details Details `json:"details,omitempty"`
}

type Details map[string]interface{}

// NewApiError create a new ApiError
func NewApiError(status int, message string, code int, details Details) *ApiError {
	return &ApiError{
		Status:  status,
		Message: message,
		Code:    code,
		Details: details,
	}
}

// Error return the error message
func (e ApiError) Error() string {
	return e.Message
}
