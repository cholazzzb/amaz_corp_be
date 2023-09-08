// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package schedulepostgres

import (
	"database/sql"

	"github.com/google/uuid"
)

type Building struct {
	ID   uuid.UUID
	Name string
}

type Friend struct {
	Member1ID uuid.UUID
	Member2ID uuid.UUID
}

type Member struct {
	ID     uuid.UUID
	UserID string
	Name   string
	Status string
	RoomID uuid.NullUUID
}

type MembersBuilding struct {
	MemberID   uuid.UUID
	BuildingID uuid.UUID
}

type Room struct {
	ID         uuid.UUID
	Name       string
	BuildingID uuid.UUID
}

type Schedule struct {
	ID     uuid.UUID
	RoomID uuid.UUID
}

type Session struct {
	ID        uuid.UUID
	RoomID    uuid.UUID
	StartTime sql.NullTime
	EndTime   sql.NullTime
}

type Task struct {
	ID           uuid.UUID
	Name         sql.NullString
	StartTime    sql.NullTime
	EndTime      sql.NullTime
	ScheduleID   uuid.UUID
	TaskDetailID uuid.UUID
}

type TaskDetail struct {
	ID         uuid.UUID
	OwnerID    uuid.NullUUID
	AssigneeID uuid.NullUUID
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
