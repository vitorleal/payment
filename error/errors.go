package errors

import (
	"net/http"
)

// InternalServerError return a new ApiError for a internal server error
func InternalServerError(err error) *g.Error {
	e := g.NewError(1000, "Internal server error", http.StatusInternalServerError)

	e.AddDetails("error", err.Error())
	return e
}

// NotFoundError return a new ApiError for a not found resource
func NotFound(err error) *g.Error {
	e := g.NewError(1001, "Not Found", http.StatusNotFound)

	e.AddDetails("error", err.Error())
	return e
}

// Unauthorized return a new ApiError for an unauthorized exception
func Unauthorized(err error) *g.Error {
	e := g.NewError(1002, "Unauthorized", http.StatusUnauthorized)

	e.AddDetails("error", err.Error())
	return e
}

// BadRequest return a new ApiError for an bad request
func BadRequest(err error) *g.Error {
	e := g.NewError(1003, "Bad request", http.StatusBadRequest)

	e.AddDetails("error", err.Error())
	return e
}
