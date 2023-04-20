package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidatorError struct {
	Message string
}

var instance = validator.New()

func Validate(i interface{}) []ValidatorError {
	var errors []ValidatorError
	err := instance.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				errors = append(errors,
					ValidatorError{
						Message: fmt.Sprintf("%s is required", err.Field()),
					},
				)
			} else {
				errors = append(errors, ValidatorError{
					Message: err.Error(),
				})
			}
		}
	}
	return errors
}
