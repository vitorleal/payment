package gateway

// Error represnets the gateway error response
type (
	Error struct {
		Status  int    `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
		Extra   Extra  `json:"extra,omitempty"`
	}

	// Extra represnets the extra error information
	Extra map[string]interface{}
)

// New create a new gateway Error
func NewError(code int, message string, status int) *Error {
	return &Error{
		Status:  status,
		Message: message,
		Code:    code,
	}
}

// Error return the error message
func (e *Error) Error() string {
	return e.Message
}

// AddDetails add extra infromation in error Extra
func (e *Error) AddDetails(extraKey string, extraValue interface{}) {
	if e.Extra == nil {
		e.Extra = Extra{}
	}

	e.Extra[extraKey] = extraValue
}
