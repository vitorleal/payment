package errors

import (
	"net/http"
)

// InternalServerError return a new ApiError for a internal server error
func InternalServerError(err error) *ApiError {
	e := NewApiError(http.StatusInternalServerError, "Internal server error", 1000, nil)

	e.Details = Details{
		"error": err.Error(),
	}

	return e
}

// NotFoundError return a new ApiError for a not found resource
func NotFound(err error) *ApiError {
	e := NewApiError(http.StatusNotFound, "Not Found", 1001, nil)

	e.Details = Details{
		"error": err.Error(),
	}

	return e
}

// Unauthorized return a new ApiError for an unauthorized exception
func Unauthorized(err error) *ApiError {
	e := NewApiError(http.StatusUnauthorized, "Unauthorized", 1002, nil)

	e.Details = Details{
		"error": err.Error(),
	}

	return e
}

// BadRequest return a new ApiError for an bad request
func BadRequest(err error) *ApiError {
	e := NewApiError(http.StatusBadRequest, "Bad request", 1003, nil)

	e.Details = Details{
		"error": err.Error(),
	}

	return e
}
