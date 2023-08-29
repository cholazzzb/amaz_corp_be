// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package locationpostgres

import (
	"database/sql"

	"github.com/google/uuid"
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

type Schedule struct {
	ID     uuid.UUID
	RoomID string
}

type Session struct {
	ID        string
	RoomID    string
	StartTime sql.NullTime
	EndTime   sql.NullTime
}

type Task struct {
	ID           uuid.UUID
	ScheduleID   uuid.UUID
	StartTime    sql.NullTime
	DurationDay  sql.NullInt32
	TaskDetailID uuid.UUID
}

type TaskDetail struct {
	ID         uuid.UUID
	Name       sql.NullString
	OwnerID    sql.NullString
	AssigneeID sql.NullString
	Status     sql.NullString
}

type TasksDependency struct {
	TaskID         uuid.NullUUID
	DependedTaskID uuid.NullUUID
}

type User struct {
	ID       string
	Username string
	Password string
	Salt     string
}
