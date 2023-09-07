package location

type BuildingQuery struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RoomQuery struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
	UserID string `json:"userID"`
	Name   string `json:"name"`
	Status string `json:"status"`
	RoomID string `json:"roomID"`
}

type JoinBuildingCommand struct {
	Name       string `json:"name" validate:"required"`
	BuildingId string `json:"buildingID" validate:"required"`
}
