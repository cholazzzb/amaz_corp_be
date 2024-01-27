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

type UserQuery struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserCommand struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	ProductID int32  `json:"productID"`
}

type ProductQuery struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type FeatureQuery struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
