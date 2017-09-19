package antifraud

import (
	"github.com/ingresse/payment/antifraud/siftscience"
	"strings"
)

type Service string

func (s Service) String() string {
	return string(s)
}

var availableServices = [...]string{
	siftscience.Name,
}

// Check if antifraud service exist
func (s Service) IsValid() bool {
	for _, service := range availableServices {
		if strings.ToLower(service) == strings.ToLower(s.String()) {
			return true
		}
	}

	return false
}
