package route_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	entLocation "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/parser"
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

func (st ScheduleTester) testTaskDependencyAfterLogin() string {
	schIDByte := tester.NewMockTest().
		Desc("/schedules should success create new schedule").
		POST(BASE_URL+"/schedules").
		Body(ent.ScheduleCommand{
			Name:   "Mock Schedule",
			RoomID: st.roomID,
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)

	scheduleID := parser.ParseResp[ent.ScheduleCommandRes](schIDByte).Data.ScheduleID

	// Need to create at least many tasks first
	// struct is currently not supported in tester module
	newTasks := []map[string]interface{}{
		{"ScheduleID": scheduleID, "Name": "Task 1", "StartTime": time.Now().AddDate(0, 0, -10).Format(time.RFC1123)},
		{"ScheduleID": scheduleID, "Name": "Task 2", "StartTime": time.Now().AddDate(0, 0, -10).Format(time.RFC1123)},
		{"ScheduleID": scheduleID, "Name": "Task 3", "StartTime": time.Now().AddDate(0, 0, -10).Format(time.RFC1123)},
		{"ScheduleID": scheduleID, "Name": "Task 4", "StartTime": time.Now().AddDate(0, 0, -10).Format(time.RFC1123)},
	}

	for _, task := range newTasks {
		tester.NewMockTest().
			Desc("/schedules/tasks should success create new task").
			POST(BASE_URL+"/tasks").
			Body(task).
			Expected(200, "").
			BuildRequest().
			WithBearer(st.bearerToken).
			Test(st.testApp, st.t)
	}

	queryParams := url.Values{}
	queryParams.Add("start-time", time.Now().AddDate(0, 0, -30).Format(time.RFC1123))
	queryParams.Add("end-time", time.Now().Format(time.RFC1123))
	fullURL := fmt.Sprintf(
		"%s?%s",
		fmt.Sprintf("%s/schedules/%s/tasks", BASE_URL, scheduleID),
		queryParams.Encode(),
	)
	getListTaskResByte := tester.NewMockTest().
		Desc("").
		GET(fullURL).
		Expected(200, "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)

	getListTaskRes := parser.ParseResp[[]ent.TaskQuery](getListTaskResByte)
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
			Expected(200, "").
			BuildRequest().
			WithBearer(st.bearerToken).
			Test(st.testApp, st.t)
	}

	return scheduleID
}

func (st ScheduleTester) testAutoSchedule(autoScheduleID string) {
	autoScheduleResByte := tester.NewMockTest().
		Desc("/schedules/:scheduleID/auto/preview show the right schedule").
		GET(BASE_URL+fmt.Sprintf("/schedules/%s/auto/preview", autoScheduleID)).
		Expected(200, "").
		BuildRequest().
		WithBearer(st.bearerToken).
		Test(st.testApp, st.t)

	autoScheduleRes := parser.ParseResp[[]ent.TaskWithDetailQuery](autoScheduleResByte)
	expected := []ent.TaskWithDetailQuery{}

	for idx := 1; idx < 5; idx++ {
		expected = append(expected, ent.TaskWithDetailQuery{
			ScheduleID: st.scheduleID,
			Name:       fmt.Sprintf("Task %d", idx),
		})
	}

	assert.Equal(st.t, len(expected), len(autoScheduleRes.Data), "num of data should be same")
	for taskIdx, real := range autoScheduleRes.Data {
		fmt.Println("REAL", taskIdx, real.Name)
		assert.Equal(st.t, expected[taskIdx].Name, real.Name, "name of data should be same")
	}
}

func TestSheduleRouteAfterLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping schedule route test in short mode.")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("--- TESTING SCHEDULE ROUTES ---")
	fmt.Println("")
	fmt.Println("")

	username := random.RandomString(12)
	_ = username + "_member"

	testApp := tester.NewMockApp().Setup("../../../.env.test")
	userID, err := tester.Register(testApp, username)
	if (err) != nil {
		panic(err)
	}
	bearerToken := tester.Login(testApp, username)

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
	myInv := parser.ParseResp[[]entLocation.BuildingMemberQuery](myInvByte)
	memberID := myInv.Data[0].MemberID

	tester.NewMockTest().
		Desc("/buildings/join should success joining member to a building").
		POST(BASE_URL+"/buildings/join").
		Body(map[string]interface{}{
			"memberID":   memberID,
			"buildingId": "bc133e57-df08-407e-b1e5-8e10c653ad3c",
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getRoomsResByte := tester.NewMockTest().
		Desc("/:buildingId/rooms should return all rooms").
		GET(BASE_URL+"/buildings/bc133e57-df08-407e-b1e5-8e10c653ad3c/rooms").
		Expected(200, "").
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
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	getListScheduleResByte := tester.NewMockTest().
		Desc("/schedules/rooms/:roomID should success return the scheduleID").
		GET(BASE_URL+"/schedules/rooms/"+getRoomRes.Rooms[0].Id).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)
	schs := parser.ParseResp[[]ent.ScheduleQuery](getListScheduleResByte)
	scheduleID := schs.Data[len(schs.Data)-1].ID

	tester.NewMockTest().
		Desc("/schedules/tasks should success create new task").
		POST(BASE_URL+"/tasks").
		Body(map[string]interface{}{
			"ScheduleID":  scheduleID,
			"StartTime":   time.Now().Format(time.RFC1123),
			"DurationDay": 3,
			"Name":        "task test 1",
		}).
		Expected(200, "").
		BuildRequest().
		WithBearer(bearerToken).
		Test(testApp, t)

	queryParams := url.Values{}
	queryParams.Add("start-time", time.Now().AddDate(0, 0, -30).Format(time.RFC1123))
	queryParams.Add("end-time", time.Now().Format(time.RFC1123))
	fullURL := fmt.Sprintf(
		"%s?%s",
		fmt.Sprintf("%s/schedules/%s/tasks", BASE_URL, scheduleID),
		queryParams.Encode(),
	)

	tester.NewMockTest().
		Desc("/schedules/:scheduleID/tasks should successfully return list task").
		GET(fullURL).
		Expected(200, "").
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

	autoScheduleID := st.testTaskDependencyAfterLogin()
	st.testAutoSchedule(autoScheduleID)
}
