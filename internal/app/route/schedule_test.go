package route_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	entLocation "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

func TestSheduleRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping schedule route test in short mode.")
	}

	newMemberName := random.RandomString(12)

	testApp := tester.NewMockApp().Setup("../../../.env.test")
	tester.Register(testApp)
	bearerToken := tester.Login(testApp)

	createMemberResByte := tester.NewMockTest().
		Desc("/members should success create member").
		POST().
		Route(BASE_URL+"/members").
		Body(map[string]interface{}{
			"name": newMemberName,
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

	getRoomsResByte := tester.NewMockTest().
		Desc("/:buildingId/rooms should return all rooms").
		GET().
		Route(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/rooms").
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
		POST().
		Body(ent.ScheduleCommand{
			RoomID: getRoomRes.Rooms[0].Id,
		}).
		Route(BASE_URL+"/schedules").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getScheduleIDResByte := tester.NewMockTest().
		Desc("/schedules/rooms/:roomID should success return the scheduleID").
		GET().
		Route(BASE_URL+"/schedules/rooms/"+getRoomRes.Rooms[0].Id).
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
		POST().
		Body(map[string]interface{}{
			"ScheduleID":  getScheduleIDRes.ScheduleID,
			"StartTime":   time.Now(),
			"DurationDay": 3,
			"Name":        "task test 1",
		}).
		Route(BASE_URL+"/tasks").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	tester.NewMockTest().
		Desc("/schedules/:scheduleID/tasks should successfully return list task").
		GET().
		Route(fmt.Sprintf("%s/schedules/%s/tasks", BASE_URL, getScheduleIDRes.ScheduleID)).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
}
