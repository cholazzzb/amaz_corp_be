package location

type Building struct {
	Id   int64
	Name string
}

type Room struct {
	Id   int64
	Name string
}

type Member struct {
	Name   string
	Status string
	UserId int64
}
