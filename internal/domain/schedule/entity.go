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

type TaskWithDetailCommandRequest struct {
	ScheduleID  string `json:"scheduleID" validate:"required"`
	StartTime   string `json:"startTime"`
	DurationDay int32  `json:"durationDay"`
	Name        string `json:"name" validate:"required"`
	OwnerID     string `json:"ownerID"`
	AssigneeID  string `json:"assigneeID"`
	Status      string `json:"status"`
}

type TaskWithDetailCommand struct {
	ScheduleID  string    `json:"scheduleID" validate:"required"`
	StartTime   time.Time `json:"startTime"`
	DurationDay int32     `json:"durationDay"`
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
	Dependencies []string  `json:"dependencies"`
}

type TaskWithDetailDependencyQuery struct {
	TaskWithDetailQuery
	Dependencies []string
}

type TaskDependencyCommand struct {
	TaskID       string `json:"taskID"`
	DependencyID string `json:"dependencyID"`
}

type TaskDependencyQuery struct {
	TaskID       string `json:"taskID"`
	DependencyID string `json:"dependencyID"`
}

type TasksDependency map[string][]string

// Interesting case to use redis
// Save this params as a key in redis
type AutoSchedulePreviewQueryParams struct {
	// ScheduleID string `json:scheduleID` params
	StartTime      string   `query:"start-time"` // mandatory
	ExcludeUsersID []string `query:"exclude-users-id"`
	ExcludeTasksID []string `query:"exclude-tasks-id"`
}
type AutoSchedulePreviewQuery struct {
	ScheduleID     string                `json:"scheduleID"`
	TasksPreview   []TaskWithDetailQuery `json:"tasksPreview"`
	AutoScheduleID string                `json:"autoScheduleID"`
}

// Interesting if use SSE....
type AutoScheduleCommand struct {
	AutoScheduleID string `json:"autoScheduleID"`
}

// this will implement bulk update to database!
