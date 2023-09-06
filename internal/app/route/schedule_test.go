package route_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	entLocation "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

func TestSheduleRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping schedule route test in short mode.")
	}

	username := random.RandomString(12)
	memberName := username + "_smember"

	testApp := tester.NewMockApp().Setup("../../../.env.test")
	tester.Register(testApp, username)
	bearerToken := tester.Login(testApp, username)

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

	getRoomsResByte := tester.NewMockTest().
		Desc("/:buildingId/rooms should return all rooms").
		GET(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/rooms").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	type GetRoomsRes struct {
		Message string             `json:"message"`
		Rooms   []entLocation.Room `json:"rooms"`
	}
	getRoomRes := GetRoomsRes{}
	json.Unmarshal(getRoomsResByte, &getRoomRes)

	tester.NewMockTest().
		Desc("/schedules should success create new schedule").
		POST(BASE_URL+"/schedules").
		Body(ent.ScheduleCommand{
			RoomID: getRoomRes.Rooms[0].Id,
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getScheduleIDResByte := tester.NewMockTest().
		Desc("/schedules/rooms/:roomID should success return the scheduleID").
		GET(BASE_URL+"/schedules/rooms/"+getRoomRes.Rooms[0].Id).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
	type GetScheduleIDRes struct {
		message    string
		ScheduleID string `json:"schedule_id"`
	}
	getScheduleIDRes := GetScheduleIDRes{}
	json.Unmarshal(getScheduleIDResByte, &getScheduleIDRes)

	tester.NewMockTest().
		Desc("/schedules/tasks should success create new task").
		POST(BASE_URL+"/tasks").
		Body(map[string]interface{}{
			"ScheduleID":  getScheduleIDRes.ScheduleID,
			"StartTime":   time.Now(),
			"DurationDay": 3,
			"Name":        "task test 1",
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/schedules/:scheduleID/tasks should successfully return list task").
		GET(fmt.Sprintf("%s/schedules/%s/tasks", BASE_URL, getScheduleIDRes.ScheduleID)).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
}
