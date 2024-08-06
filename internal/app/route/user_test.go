package route_test

import (
	"encoding/json"
	"testing"

	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

const BASE_URL = "http://localhost:8080/api/v1"

func TestUserRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping user route test in short mode.")
	}

	testApp := tester.NewMockApp().Setup("../../../.env.test")

	tester.NewMockTest().
		Desc("unauthorized get api").
		GET(BASE_URL+"/users").
		Expected(401, "").
		BuildRequest().
		Test(testApp, t)

	tester.NewMockTest().
		Desc("get HTTP status 404, when route is not exists").
		GET(BASE_URL+"/not-found").
		Expected(404, "").
		BuildRequest().
		Test(testApp, t)

	// TODO: Fix login api, when user not found return 400
	//  tester.NewMockTest().
	// 	Desc("login").
	// 	POST().
	// 	Route(BASE_URL+"/login").
	// 	Body(map[string]interface{}{
	// 		"username": "gaada",
	// 		"password": "gaada",
	// 	}).
	// 	Expected(400, "", "").
	//  BuildRequest2().
	//  Test(testApp, t)
}

func TestUserRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping user route test in short mode.")
	}

	newUsername := random.RandomString(12)

	testApp := tester.NewMockApp().Setup("../../../.env.test")

	tester.NewMockTest().
		Desc("/register when user not register should successfull").
		POST(BASE_URL+"/register").
		Body(map[string]interface{}{
			"username": newUsername,
			"password": newUsername,
		}).
		Expected(200, "").
		BuildRequest().
		Test(testApp, t)

	loginResByte := tester.NewMockTest().
		Desc("/login after register should successfull").
		POST(BASE_URL+"/login").
		Body(map[string]interface{}{
			"username": newUsername,
			"password": newUsername,
		}).
		Expected(200, "").
		BuildRequest().
		Test(testApp, t)

	type LoginRes struct {
		Token string `json:"token"`
	}
	loginRes := LoginRes{}
	json.Unmarshal(loginResByte, &loginRes)

}
