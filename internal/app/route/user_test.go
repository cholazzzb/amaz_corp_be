package route_test

import (
	"testing"

	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

const BASE_URL = "http://localhost:8080/api/v1"

func TestUserRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping user route test in short mode.")
	}

	tests := tester.MockTester{}

	test1 := *tester.NewMockTest().
		Desc("unauthorized get api").
		GET().
		Route(BASE_URL+"/users").
		Expected(401, "", "")
	tests.AddTest(test1)

	test2 := *tester.NewMockTest().
		Desc("get HTTP status 404, when route is not exists").
		GET().
		Route(BASE_URL+"/not-found").
		Expected(404, "", "")
	tests.AddTest(test2)

	// TODO: Fix login api, when user not found return 400
	// test3 := *tester.NewMockTest().
	// 	Desc("login").
	// 	POST().
	// 	Route(BASE_URL+"/login").
	// 	Body(map[string]interface{}{
	// 		"username": "gaada",
	// 		"password": "gaada",
	// 	}).
	// 	Expected(400, "", "")
	// tests.AddTest(test3)

	tests.Setup("../../../.env.dev")
	tests.Test(t)
}

func TestAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping user route test in short mode.")
	}

	newUsername := random.RandomString(12)

	tests := tester.MockTester{}

	test1 := *tester.NewMockTest().
		Desc("/register when user not register should successfull").
		POST().
		Route(BASE_URL+"/register").
		Body(map[string]interface{}{
			"username": newUsername,
			"password": newUsername,
		}).
		Expected(200, "", "").
		WithAuth()
	tests.AddTest(test1)

	test2 := *tester.NewMockTest().
		Desc("/login after register should successfull").
		POST().
		Route(BASE_URL+"/login").
		Body(map[string]interface{}{
			"username": newUsername,
			"password": newUsername,
		}).
		Expected(200, "", "").
		WithAuth()
	tests.AddTest(test2)

	test3 := *tester.NewMockTest().
		Desc("/members").
		POST().
		Route(BASE_URL+"/members").
		Body(map[string]interface{}{
			"name": newUsername + "_name",
		}).
		Expected(200, "", "").
		WithAuth()
	tests.AddTest(test3)

	tests.Setup("../../../.env.dev")
	tests.Test(t)

}
