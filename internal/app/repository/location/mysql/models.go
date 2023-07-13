// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package location

import (
	"database/sql"
)

type Building struct {
	ID   string
	Name string
}

type Friend struct {
	Member1ID string
	Member2ID string
}

type Member struct {
	ID     string
	UserID string
	Name   string
	Status string
	RoomID sql.NullString
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
	StartTime sql.NullTime
	EndTime   sql.NullTime
}

type User struct {
	ID       string
	Username string
	Password string
	Salt     string
}
