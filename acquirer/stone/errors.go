package stone

import (
	g "github.com/ingresse/payment/gateway"
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

// BadRequestError return a new badRequest error
// with custom message, code and extra informations
func BadRequestError(message string, report *ErrorReport, code int) *g.Error {
	e := g.NewError(code, message, http.StatusBadRequest)

	for _, r := range report.ErrorItemCollection {
		if r.ErrorField != "" {
			e.AddDetails("field", r.ErrorField)
		}

		e.AddDetails("info", r.Description)
		e.AddDetails("severity", r.SeverityCode)
	}

	return e
}
