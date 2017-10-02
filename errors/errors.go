package errors

import (
	"net/http"
)

// InternalServerError return a new ApiError for a internal server error
func InternalServerError(err error) *ApiError {
	e := NewApiError(1000, "Internal server error", http.StatusInternalServerError, nil)

	e.AddDetails("error", err.Error())
	return e
}

// NotFoundError return a new ApiError for a not found resource
func NotFound(err error) *ApiError {
	e := NewApiError(1001, "Not Found", http.StatusNotFound, nil)

	e.AddDetails("error", err.Error())
	return e
}

// Unauthorized return a new ApiError for an unauthorized exception
func Unauthorized(err error) *ApiError {
	e := NewApiError(1002, "Unauthorized", http.StatusUnauthorized, nil)

	e.AddDetails("error", err.Error())
	return e
}

// BadRequest return a new ApiError for an bad request
func BadRequest(err error) *ApiError {
	e := NewApiError(1003, "Bad request", http.StatusBadRequest, nil)

	e.AddDetails("error", err.Error())
	return e
}
