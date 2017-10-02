package stone

import (
	"github.com/ingresse/payment/errors"
	"net/http"
)

const (
	AuthorizeError     string = "Error authorizing the sale"
	AuthorizeErrorCode int    = 2000

	CaptureError     string = "Error capturing the sale"
	CaptureErrorCode int    = 2001

	GetSaleError     string = "Error requesting for the sale data"
	GetSaleErrorCode int    = 2002

	CancelError     string = "Error canceling the sale"
	CancelErrorCode int    = 2003
)

// BadRequestError return a new badRequest error with custom message and code
func BadRequestError(message string, report *ErrorReport, code int) *errors.ApiError {
	e := errors.NewApiError(code, message, http.StatusBadRequest, nil)

	for _, reportError := range report.ErrorItemCollection {
		e.AddDetails("details", reportError.Description)
	}

	return e
}
