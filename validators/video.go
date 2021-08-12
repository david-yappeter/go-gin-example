package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	fmt.Printf("%+v\n", field.Field())
	return strings.Contains(field.Field().String(), "Cool")
}
