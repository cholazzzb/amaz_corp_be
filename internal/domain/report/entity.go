package report

import "github.com/cholazzzb/amaz_corp_be/internal/domain/location"

type ReportByScheduleQuery struct {
	ScheduleID         string                     `json:"scheduleID"`
	TotalTask          int32                      `json:"totalTask"`
	AvgTaskDuration    int32                      `json:"avgTaskDuration"`
	MedianTaskDuration int32                      `json:"medianTaskDuration"`
	TaskSortByAssignee []string                   `json:"taskSortByAssignee"`
	AssigneeMap        map[string]location.Member `json:"assigneeMap"`
}
