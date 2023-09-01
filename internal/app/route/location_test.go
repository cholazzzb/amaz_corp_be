package route_test

import (
	"encoding/json"
	"testing"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

func TestLocationRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping location route test in short mode.")
	}

	testApp := tester.NewMockApp().Setup("../../../.env.test")
	bearerToken := tester.Login(testApp)

	tester.NewMockTest().
		Desc("/buildings should return all buildings of current member").
		GET().
		Route(BASE_URL+"/buildings").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/buildings/all should return all buildings").
		GET().
		Route(BASE_URL+"/buildings/all").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	// Note: buildingID from the seeder
	tester.NewMockTest().
		Desc("/buildings/:buildingId/rooms should return all rooms in the building").
		GET().
		Route(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/rooms").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/buildings/:buildingID/members should return all members in the building").
		GET().
		Route(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/members").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	createMemberResByte := tester.NewMockTest().
		Desc("/members should success create member").
		POST().
		Route(BASE_URL+"/members").
		Body(map[string]interface{}{
			"name": "testing1",
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	type CreateMemberRes struct {
		Member user.Member `json:"member"`
	}
	createMemberRes := CreateMemberRes{}
	json.Unmarshal(createMemberResByte, &createMemberRes)

	// Note: buildingID from the seeder
	tester.NewMockTest().
		Desc("/buildings/join should success joining member to a building").
		POST().
		Body(map[string]interface{}{
			"memberId":   createMemberRes.Member.ID,
			"buildingId": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Route(BASE_URL+"/buildings/join").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/buildings/leave should success").
		DELETE(BASE_URL+"/buildings/leave").
		Body(map[string]interface{}{
			"MemberId":   createMemberRes.Member.ID,
			"BuildingId": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
}
