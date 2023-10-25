package schedule

import (
	"time"
)

type ScheduleCommand struct {
	Name   string `json:"name" validate:"required"`
	RoomID string `json:"roomID" validate:"required"`
}

type ScheduleQuery struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	RoomID string `json:"roomID"`
}

type TaskCommand struct {
	ScheduleID   string
	StartTime    time.Time
	DurationDay  int32
	TaskDetailID string
}

type TaskQuery struct {
	ID           string    `json:"ID"`
	Name         string    `json:"name"`
	StartTime    time.Time `json:"startTime"`
	DurationDay  int32     `json:"durationDay"`
	ScheduleID   string    `json:"scheduleID"`
	EndTime      time.Time `json:"endTime"`
	TaskDetailID string    `json:"taskDetailID"`
}

type TaskQueryFilter struct {
	StartTime *time.Time
	EndTime   *time.Time
}

type TaskQueryFilterParams struct {
	StartTime string `query:"start-time"`
	EndTime   string `query:"end-time"`
}

type TaskDetailCommand struct {
	Name       string
	OwnerID    string
	AssigneeID string
	Status     string
}

type TaskDetailQuery struct {
	ID         string `json:"ID"`
	OwnerID    string `json:"ownerID"`
	AssigneeID string `json:"assigneeID"`
	Status     string `json:"status"`
}

type TaskWithDetailCommand struct {
	ScheduleID  string    `json:"scheduleID" validate:"required"`
	StartTime   time.Time `json:"startTime" validate:"required"`
	DurationDay int32     `json:"durationDay" validate:"required,gte=0,lte=14"`
	Name        string    `json:"name" validate:"required"`
	OwnerID     string    `json:"ownerID"`
	AssigneeID  string    `json:"assigneeID"`
	Status      string    `json:"status"`
}

type TaskWithDetailQuery struct {
	TaskID       string    `json:"taskID"`
	ScheduleID   string    `json:"scheduleID"`
	StartTime    time.Time `json:"startTime"`
	DurationDay  int32     `json:"durationDay"`
	TaskDetailID string    `json:"taskDetailID"`
	Name         string    `json:"name"`
	OwnerID      string    `json:"ownerID"`
	AssigneeID   string    `json:"assigneeID"`
	Status       string    `json:"status"`
}
