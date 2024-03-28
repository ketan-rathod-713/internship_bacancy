package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	// if the string contains word cool then return true else return false
	return strings.Contains(field.Field)
}
