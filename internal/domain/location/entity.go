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

type MemberQuery struct {
	MemberID string
	Name     string
	Status   string
}
