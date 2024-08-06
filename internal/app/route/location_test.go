package route_test

import (
	"fmt"
	"testing"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	"github.com/cholazzzb/amaz_corp_be/pkg/parser"
	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

func TestLocationRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping location route test in short mode.")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("--- TESTING SCHEDULE ROUTES ---")
	fmt.Println("")
	fmt.Println("")

	testApp := tester.NewMockApp().Setup("../../../.env.test")

	username := random.RandomString(12)
	memberName := username + "_member"

	userID, err := tester.Register(testApp, username)
	if (err) != nil {
		panic(err)
	}
	bearerToken := tester.Login(testApp, username)

	tester.NewMockTest().
		Desc("/buildings/all should return all buildings").
		GET(BASE_URL+"/buildings/all").
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	// Note: buildingID from the seeder
	tester.NewMockTest().
		Desc("/buildings/:buildingId/rooms should return all rooms in the building").
		GET(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/rooms").
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/buildings/:buildingID/members should return all members in the building").
		GET(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/members").
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	// Note: buildingID from the seeder
	tester.NewMockTest().
		Desc("/buildings/invite should success inviting member to a building").
		POST(BASE_URL+"/buildings/invite").
		Body(map[string]interface{}{
			"userID":     userID,
			"buildingID": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	myInvByte := tester.NewMockTest().
		GET(BASE_URL+"/buildings/invitation").
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
	myInv := parser.ParseResp[[]ent.BuildingMemberQuery](myInvByte)
	memberID := myInv.Data[0].MemberID

	tester.NewMockTest().
		Desc("/buildings/join should success joining member to a building").
		POST(BASE_URL+"/buildings/join").
		Body(map[string]interface{}{
			"memberID":   memberID,
			"buildingID": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/members should success edit member name").
		PUT(BASE_URL+"/members").
		Body(map[string]interface{}{
			"memberID": memberID,
			"name":     memberName,
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getMemberByNameByte := tester.NewMockTest().
		Desc("/members/:name should return the true member").
		GET(BASE_URL+"/members?name="+memberName).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
	member := parser.ParseResp[ent.MemberQuery](getMemberByNameByte)

	tester.NewMockTest().
		Desc("/buildings/leave should success").
		DELETE(BASE_URL+"/buildings/leave").
		Body(map[string]interface{}{
			"memberID":   member.Data.ID,
			"buildingID": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
}
