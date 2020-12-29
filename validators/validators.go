package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// custom validation 구현
func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
