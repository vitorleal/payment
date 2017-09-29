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

// FormatError return a new error
func BadRequestError(message string, report *ErrorReport, code int) error {
	e := errors.NewApiError(http.StatusBadRequest, message, code, nil)

	return e
}
