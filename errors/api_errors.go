package errors

import (
	g "github.com/gin-gonic/gin"
)

type ApiError struct {
	Status  int     `json:"status,omitempty"`
	Message string  `json:"message,omitempty"`
	Code    int     `json:"code,omitempty"`
	Details Details `json:"details,omitempty"`
}

type Details map[string]interface{}

// NewApiError create a new ApiError
func NewApiError(code int, message string, status int, details Details) *ApiError {
	return &ApiError{
		Status:  status,
		Message: message,
		Code:    code,
		Details: details,
	}
}

// Error return the error message
func (e *ApiError) Error() string {
	return e.Message
}

// AddDetails add extra infromation in error Details
func (e *ApiError) AddDetails(extraKey string, extraValue interface{}) {
	if e.Details == nil {
		e.Details = Details{}
	}

	e.Details[extraKey] = extraValue
}

// Json return the error in json format
func (e *ApiError) Json() g.H {
	return g.H{
		"error":   e.Message,
		"status":  e.Status,
		"code":    e.Code,
		"details": e.Details,
	}
}
