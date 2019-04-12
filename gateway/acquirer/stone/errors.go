package stone

import (
	e "github.com/ingresse/payment/errors"
)

var (
	AuthorizeError = &e.CustomError{
		Message: "Error authorizing the sale",
		Code:    2000,
	}

	CaptureError = &e.CustomError{
		Message: "Error capturing the sale",
		Code:    2001,
	}

	GetSaleError = &e.CustomError{
		Message: "Error requesting for the sale data",
		Code:    2002,
	}

	CancelError = &e.CustomError{
		Message: "Error canceling the sale",
		Code:    2003,
	}
)

// Response return a new badRequest error
// with custom message, code and extra informations
func ResponseError(errorType *e.CustomError, report *ErrorReport) *e.Error {
	err := e.BadRequestError(errorType.Code, errorType.Message)

	for _, r := range report.ErrorItemCollection {
		if r.ErrorField != "" {
			err.AddDetails("field", r.ErrorField)
		}

		err.AddDetails("info", r.Description)
		err.AddDetails("severity", r.SeverityCode)
	}

	return err
}
