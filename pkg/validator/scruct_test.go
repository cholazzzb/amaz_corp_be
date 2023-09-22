package validator_test

import (
	"testing"

	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
	"github.com/stretchr/testify/assert"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func TestRegisterValidation(t *testing.T) {
	passed := RegisterRequest{Username: "username", Password: "password"}
	errors1 := validator.Validate(passed)
	assert.Equal(t, []validator.ValidatorError(nil), errors1, "Right struct should not get validate error")

	notPassed := RegisterRequest{}
	errors2 := validator.Validate(notPassed)
	assert.Equal(t,
		[]validator.ValidatorError{
			{Field: "Username", Message: "This field is required"},
			{Field: "Password", Message: "This field is required"},
		},
		errors2, "Wrong Struct should not passing validation",
	)
}
