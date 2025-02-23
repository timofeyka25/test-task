package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(i interface{}) error {
	err := validate.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation error: field '%s' failed on the '%s' tag", err.Field(), err.Tag())
		}
	}
	return nil
}
