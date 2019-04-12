package errors

import (
	"net/http"
)

// Error represents an error
type (
	Error struct {
		HttpStatus int    `json:"httpStatus,omitempty"`
		Message    string `json:"message,omitempty"`
		Code       int    `json:"code,omitempty"`
		Extra      Extra  `json:"extra,omitempty"`
	}

	// Extra represents the extra error information
	Extra map[string]interface{}

	// CustomErrorType represents a custom error code and error and message
	CustomError struct {
		Message string
		Code    int
	}
)

var (
	Unauthorized = &CustomError{
		Code:    1000,
		Message: "Unauthorized access",
	}

	BadRequest = &CustomError{
		Code:    1001,
		Message: "Bad request",
	}
)

// New create a new error
func NewError(custom *CustomError, httpStatus int) *Error {
	e := &Error{
		Code:    custom.Code,
		Message: custom.Message,
	}

	if httpStatus > 0 {
		e.HttpStatus = httpStatus
	}

	return e
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

// UnauthorizedError creates a new error with unauthorized http status
func UnauthorizedError(code int, message string) *Error {
	err := NewError(BadRequest, http.StatusUnauthorized)

	if code > 0 {
		err.Code = code
	}

	if message != "" {
		err.Message = message
	}

	return err
}

// BadRequestError creates a new error with bad request http status
func BadRequestError(code int, message string) *Error {
	err := NewError(BadRequest, http.StatusBadRequest)

	if code > 0 {
		err.Code = code
	}

	if message != "" {
		err.Message = message
	}

	return err
}
