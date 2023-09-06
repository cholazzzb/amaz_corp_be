package location

type Building struct {
	Id   string
	Name string
}

type Room struct {
	Id   string
	Name string
}

type Member struct {
	Name   string
	Status string
	UserId string
}

type MemberCommand struct {
	UserID string
	Name   string
	Status string
}

type MemberQuery struct {
	ID     string `json:"id"`
	UserID string
	Name   string
	Status string
	RoomID string
}

type JoinBuildingCommand struct {
	Name       string `json:"name" validate:"required"`
	BuildingId string `json:"buildingId" validate:"required"`
}
