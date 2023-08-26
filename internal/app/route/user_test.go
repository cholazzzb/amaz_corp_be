package route_test

import (
	"encoding/json"
	"testing"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
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
		GET().
		Route(BASE_URL+"/users").
		Expected(401, "", "").
		BuildRequest().
		Test(testApp, t)

	tester.NewMockTest().
		Desc("get HTTP status 404, when route is not exists").
		GET().
		Route(BASE_URL+"/not-found").
		Expected(404, "", "").
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
		POST().
		Route(BASE_URL+"/register").
		Body(map[string]interface{}{
			"username": newUsername,
			"password": newUsername,
		}).
		Expected(200, "", "").
		BuildRequest().
		Test(testApp, t)

	loginResByte := tester.NewMockTest().
		Desc("/login after register should successfull").
		POST().
		Route(BASE_URL+"/login").
		Body(map[string]interface{}{
			"username": newUsername,
			"password": newUsername,
		}).
		Expected(200, "", "").
		BuildRequest().
		Test(testApp, t)

	type LoginRes struct {
		Token string `json:"token"`
	}
	loginRes := LoginRes{}
	json.Unmarshal(loginResByte, &loginRes)

	createMemberResByte := tester.NewMockTest().
		Desc("/members should success create member").
		POST().
		Route(BASE_URL+"/members").
		Body(map[string]interface{}{
			"name": newUsername,
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(loginRes.Token).
		Test(testApp, t)

	type CreateMemberRes struct {
		Member user.Member `json:"member"`
	}
	createMemberRes := CreateMemberRes{}
	json.Unmarshal(createMemberResByte, &createMemberRes)
}
