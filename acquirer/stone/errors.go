package stone

import (
	g "github.com/ingresse/payment/gateway"
	"net/http"
)

type ErrorType struct {
	Message string
	Code    int
}

var (
	AuthorizeError = &ErrorType{
		Message: "Error authorizing the sale",
		Code:    2000,
	}

	CaptureError = &ErrorType{
		Message: "Error capturing the sale",
		Code:    2001,
	}

	GetSaleError = &ErrorType{
		Message: "Error requesting for the sale data",
		Code:    2002,
	}

	CancelError = &ErrorType{
		Message: "Error canceling the sale",
		Code:    2003,
	}
)

// BadRequestError return a new badRequest error
// with custom message, code and extra informations
func BadRequestError(errorType *ErrorType, report *ErrorReport) *g.Error {
	e := g.NewError(errorType.Code, errorType.Message, http.StatusBadRequest)

	for _, r := range report.ErrorItemCollection {
		if r.ErrorField != "" {
			e.AddDetails("field", r.ErrorField)
		}

		e.AddDetails("info", r.Description)
		e.AddDetails("severity", r.SeverityCode)
	}

	return e
}
