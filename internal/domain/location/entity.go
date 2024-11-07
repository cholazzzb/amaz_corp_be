package location

type BuildingCommand struct {
	Name string `json:"name"`
}

type BuildingQuery struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BuildingMemberQuery struct {
	BuildingID   string `json:"buildingID"`
	BuildingName string `json:"buildingName"`
	MemberID     string `json:"memberID"`
}

type RoomCommand struct {
	Name       string `json:"name" validate:"required"`
	BuildingID string `json:"buildingID" validate:"required"`
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

type InviteMemberToBuildingCommand struct {
	UserID     string `json:"userID" validate:"required,min=36,max=36"`
	BuildingID string `json:"buildingID" validate:"required,min=36,max=36"`
}

type JoinBuildingCommand struct {
	MemberID   string `json:"memberID" validate:"required"`
	BuildingID string `json:"buildingID" validate:"required,min=36,max=36"`
}

type RenameMemberCommand struct {
	MemberID string `json:"memberID" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
