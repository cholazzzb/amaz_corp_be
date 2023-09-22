package validator

import (
	"errors"

	v10 "github.com/go-playground/validator/v10"
)

type ValidatorError struct {
	Field   string
	Message string
}

var instance = v10.New()

func Validate(i interface{}) []ValidatorError {
	var errs []ValidatorError
	err := instance.Struct(i)
	if err != nil {
		var ve v10.ValidationErrors
		if errors.As(err, &ve) {
			for _, err := range ve {
				errs = append(errs, ValidatorError{
					err.Field(),
					msgForTag(err.Tag(), err.Error()),
				})
			}
		} else {
			errs = append(errs, ValidatorError{
				"Some field are missing",
				"Check the api contract!",
			})
		}
	}
	return errs
}

func msgForTag(tag, err string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return err
}
