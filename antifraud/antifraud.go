package antifraud

import (
	"github.com/ingresse/payment/antifraud/siftscience"
	"strings"
)

type Service string

// String return the service represnete in string
func (s Service) String() string {
	return string(s)
}

// List of availabel antifraud services
var availableServices = [...]string{
	siftscience.Name,
}

// IsValid validate that the antifraud service is implemented
func (s Service) IsValid() bool {
	for _, service := range availableServices {
		if strings.ToLower(service) == strings.ToLower(s.String()) {
			return true
		}
	}

	return false
}
