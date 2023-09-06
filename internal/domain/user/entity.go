package user

type Building struct {
	ID   string
	Name string
}

type Friend struct {
	Member1ID string
	Member2ID string
}

type MembersBuilding struct {
	MemberID   string
	BuildingID string
}

type Room struct {
	ID         string
	Name       string
	BuildingID string
}

type Session struct {
	ID        string
	RoomID    string
	StartTime interface{}
	EndTime   interface{}
}

type User struct {
	ID       string
	Username string
	Password string
	Salt     string
}
