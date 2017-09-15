package siftscience

import (
	"strings"
)

const OK = "OK"

// Siftscience score
type Score struct {
	Status       uint32  `json:"status,omitempty"`
	ErrorMessage string  `json:"error_message,omitempty"`
	UserId       string  `json:"user_id,omitempty"`
	Scores       *Scores `json:"scores,omitempty"`
}

// Siftscience scores
type Scores struct {
	PaymentAbuse *PaymentAbuse `json:"payment_abuse,omitempty"`
	AccountAbuse *AccountAbuse `json:"account_abuse,omitempty"`
}

// Siftscience payment abuse data
type PaymentAbuse struct {
	Score float64 `json:"score,omitempty"`
}

// Siftscience account abuse data
type AccountAbuse struct {
	Score float64 `json:"score,omitempty"`
}

// Check if score response is Ok
func (score *Score) IsOk() bool {
	if strings.ToUpper(score.ErrorMessage) == OK && score.Status == 0 {
		return true
	}

	return false
}
