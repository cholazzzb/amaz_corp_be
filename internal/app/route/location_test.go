package route_test

import (
	"encoding/json"
	"testing"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
	"github.com/stretchr/testify/assert"
)

func TestLocationRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping location route test in short mode.")
	}

	testApp := tester.NewMockApp().Setup("../../../.env.test")

	username := random.RandomString(12)
	memberName := username + "_member"

	tester.Register(testApp, username)
	bearerToken := tester.Login(testApp, username)

	tester.NewMockTest().
		Desc("/buildings should return all buildings of current member").
		GET(BASE_URL+"/buildings").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/buildings/all should return all buildings").
		GET(BASE_URL+"/buildings/all").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	// Note: buildingID from the seeder
	tester.NewMockTest().
		Desc("/buildings/:buildingId/rooms should return all rooms in the building").
		GET(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/rooms").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/buildings/:buildingID/members should return all members in the building").
		GET(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/members").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	// Note: buildingID from the seeder
	tester.NewMockTest().
		Desc("/buildings/join should success joining member to a building").
		POST(BASE_URL+"/buildings/join").
		Body(map[string]interface{}{
			"name":       memberName,
			"buildingId": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getMemberByNameByte := tester.NewMockTest().
		Desc("/members/:name should return the true member").
		GET(BASE_URL+"/members/name/"+memberName+"/search").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	type GetMemberByNameRes struct {
		Message string          `json:"message"`
		Member  ent.MemberQuery `json:"member"`
	}
	getMemberByNameRes := GetMemberByNameRes{}
	json.Unmarshal(getMemberByNameByte, &getMemberByNameRes)
	assert.Equalf(t, memberName, getMemberByNameRes.Member.Name, "the respond memnber name should be same with request")

	tester.NewMockTest().
		Desc("/buildings/leave should success").
		DELETE(BASE_URL+"/buildings/leave").
		Body(map[string]interface{}{
			"memberID":   getMemberByNameRes.Member.ID,
			"buildingID": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
}
