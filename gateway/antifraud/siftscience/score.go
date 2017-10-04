package siftscience

import (
	"strings"
)

const ErrorOK = "OK"
const SuccessStatus = 0

type (
	// Score represents the Siftscience score data
	Score struct {
		Status       uint32  `json:"status,omitempty"`
		ErrorMessage string  `json:"error_message,omitempty"`
		UserId       string  `json:"user_id,omitempty"`
		Scores       *Scores `json:"scores,omitempty"`
	}

	// Scores represents the Siftscience scores types
	Scores struct {
		PaymentAbuse *PaymentAbuse `json:"payment_abuse,omitempty"`
		AccountAbuse *AccountAbuse `json:"account_abuse,omitempty"`
	}

	// PaymentAbuse represnet Siftscience payment abuse score data
	PaymentAbuse struct {
		Score float64 `json:"score,omitempty"`
	}

	// AccountAbuse represetns Siftscience account abuse score data
	AccountAbuse struct {
		Score float64 `json:"score,omitempty"`
	}
)

// IsOk verify if Siftscience score api response is Ok
func (score *Score) IsOk() bool {
	if strings.ToUpper(score.ErrorMessage) == ErrorOK && score.Status == SuccessStatus {
		return true
	}

	return false
}
