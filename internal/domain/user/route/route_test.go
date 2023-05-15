package user_test

import (
	"testing"

	testhelper "github.com/cholazzzb/amaz_corp_be/pkg/test_helper"
)

func TestUserRoute(t *testing.T) {
	tests := []testhelper.TestRoute{
		{
			Method:      "GET",
			Description: "register",
			Route:       "/api/v1/login",
			Body: map[string]string{
				"username": "test1",
				"password": "password1",
			},
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "OK",
		},
	}

	app := testhelper.Setup(t)

	testhelper.RunTests(t, app, tests)
}
