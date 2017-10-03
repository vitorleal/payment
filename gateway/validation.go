package gateway

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/ingresse/payment/antifraud/siftscience"
	validator "gopkg.in/go-playground/validator.v8"
	"reflect"
	"strings"
)

// Register custom validations
func RegisterValidations() {
	binding.Validator.RegisterValidation("creditCardBrand", ValidCreditCards)
	binding.Validator.RegisterValidation("antifraud", ValidAntifraud)
	binding.Validator.RegisterValidation("acquirer", ValidAcquirer)
}

// ValidCreditCards validate the available creditcards
func ValidCreditCards(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	cards := [...]string{
		"visa", "mastercard", "amex", "elo",
	}

	for _, card := range cards {
		if strings.ToLower(field.String()) == strings.ToLower(card) {
			return true
		}
	}

	return false
}

// ValidAcquirer validate the availables acquirers
func ValidAcquirer(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	acquirers := [...]string{
		"stone",
	}

	for _, acq := range acquirers {
		if strings.ToLower(field.String()) == strings.ToLower(acq) {
			return true
		}
	}

	return false
}

// ValidAntifraud validate the available antifrauds services
func ValidAntifraud(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	services := [...]string{
		siftscience.Name,
	}

	for _, service := range services {
		fmt.Printf("%s", service)
		fmt.Printf("%s", field.String())

		if strings.ToLower(field.String()) == strings.ToLower(service) {
			return true
		}
	}

	return false
}
