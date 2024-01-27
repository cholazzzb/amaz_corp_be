package route_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	entLocation "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/random"
	"github.com/cholazzzb/amaz_corp_be/pkg/tester"
)

type ScheduleTester struct {
	scheduleID  string
	roomID      string
	bearerToken string
	testApp     *fiber.App
	t           *testing.T
}

func (st ScheduleTester) testTaskDependencyAfterLogin() {
	tester.NewMockTest().
		Desc("/schedules should success create new schedule").
		POST(BASE_URL+"/schedules").
		Body(ent.ScheduleCommand{
			Name:   "Mock Schedule",
			RoomID: st.roomID,
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)

	getListScheduleResByte := tester.NewMockTest().
		Desc("/schedules/rooms/:roomID should success return the scheduleID").
		GET(BASE_URL+"/schedules/rooms/"+st.roomID).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)
	type GetListScheduleRes struct {
		Message string              `json:"message"`
		Data    []ent.ScheduleQuery `json:"data"`
	}
	getListScheduleRes := GetListScheduleRes{}
	json.Unmarshal(getListScheduleResByte, &getListScheduleRes)

	schs := getListScheduleRes.Data
	scheduleID := schs[len(schs)-1].ID

	// Need to create at least many tasks first
	// struct is currently not supported in tester module
	newTasks := []map[string]interface{}{
		{"ScheduleID": scheduleID, "Name": "Task 1"},
		{"ScheduleID": scheduleID, "Name": "Task 2"},
		{"ScheduleID": scheduleID, "Name": "Task 3"},
		{"ScheduleID": scheduleID, "Name": "Task 4"},
	}

	for _, task := range newTasks {
		type CreateTaskRes struct {
			Message string   `json:"message"`
			Data    []string `json:"data"`
		}
		tester.NewMockTest().
			Desc("/schedules/tasks should success create new task").
			POST(BASE_URL+"/tasks").
			Body(task).
			Expected(200, "", "").
			BuildRequest().
			WithBearer(st.bearerToken).
			Test(st.testApp, st.t)
	}

	getListTaskResByte := tester.NewMockTest().
		Desc("").
		GET(BASE_URL+"/schedules/"+scheduleID+"/tasks").
		Expected(200, "", "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)
	type GetListTaskRes struct {
		Message string          `json:"message"`
		Data    []ent.TaskQuery `json:"data"`
	}

	getListTaskRes := GetListTaskRes{}
	json.Unmarshal(getListTaskResByte, &getListTaskRes)
	assert.Equal(st.t, 4, len(getListTaskRes.Data), "all tasks should save on DB")

	for taskIdx := 1; taskIdx < 4; taskIdx++ {
		task := getListTaskRes.Data[taskIdx]
		prevTask := getListTaskRes.Data[taskIdx-1]

		tester.NewMockTest().
			Desc("/tasks/dependency/ should success connect task dependency").
			POST(BASE_URL+"/tasks/dependency").
			Body(map[string]interface{}{
				"taskID":       task.ID,
				"dependencyID": prevTask.ID,
			}).
			Expected(200, "", "").
			BuildRequest().
			WithBearer(st.bearerToken).
			Test(st.testApp, st.t)
	}
}

func (st ScheduleTester) testAutoSchedule() {
	autoScheduleResByte := tester.NewMockTest().
		Desc("/schedules/:scheduleID/auto/preview show the right schedule").
		GET(BASE_URL+fmt.Sprintf("/schedules/%s/auto/preview", st.scheduleID)).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)
	type AutoScheduleRes struct {
		Message string                    `json:"message"`
		Data    []ent.TaskWithDetailQuery `json:"data"`
	}
	autoScheduleRes := AutoScheduleRes{}
	json.Unmarshal(autoScheduleResByte, &autoScheduleRes)
	expected := []ent.TaskWithDetailQuery{}

	for idx := 1; idx < 5; idx++ {
		expected = append(expected, ent.TaskWithDetailQuery{
			ScheduleID: st.scheduleID,
			Name:       fmt.Sprintf("Task %d", idx),
		})
	}

	assert.Equal(st.t, expected, autoScheduleRes.Data, "task should sorted")
}

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
		Message string                  `json:"message"`
		Rooms   []entLocation.RoomQuery `json:"rooms"`
	}
	getRoomRes := GetRoomsRes{}
	json.Unmarshal(getRoomsResByte, &getRoomRes)

	roomID := getRoomRes.Rooms[0].Id

	tester.NewMockTest().
		Desc("/schedules should success create new schedule").
		POST(BASE_URL+"/schedules").
		Body(ent.ScheduleCommand{
			Name:   "Mock Schedule",
			RoomID: getRoomRes.Rooms[0].Id,
		}).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getListScheduleResByte := tester.NewMockTest().
		Desc("/schedules/rooms/:roomID should success return the scheduleID").
		GET(BASE_URL+"/schedules/rooms/"+getRoomRes.Rooms[0].Id).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
	type GetListScheduleRes struct {
		Message string              `json:"message"`
		Data    []ent.ScheduleQuery `json:"data"`
	}
	getListScheduleRes := GetListScheduleRes{}
	json.Unmarshal(getListScheduleResByte, &getListScheduleRes)

	schs := getListScheduleRes.Data
	scheduleID := schs[len(schs)-1].ID

	tester.NewMockTest().
		Desc("/schedules/tasks should success create new task").
		POST(BASE_URL+"/tasks").
		Body(map[string]interface{}{
			"ScheduleID":  scheduleID,
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
		GET(fmt.Sprintf("%s/schedules/%s/tasks", BASE_URL, scheduleID)).
		Expected(200, "", "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	st := ScheduleTester{
		scheduleID:  scheduleID,
		roomID:      roomID,
		bearerToken: bearerToken,
		testApp:     testApp,
		t:           t,
	}

	st.testTaskDependencyAfterLogin()
	st.testAutoSchedule()
}
