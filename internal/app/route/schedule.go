package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type ScheduleRoute struct {
	fr fiber.Router
	h  *handler.ScheduleHandler
}

func NewScheduleRoute(fr fiber.Router, h *handler.ScheduleHandler) *ScheduleRoute {
	return &ScheduleRoute{
		fr, h,
	}
}

func (r *ScheduleRoute) InitRoute(am middleware.Middleware) {
	scheduleApi := r.fr.Group("/schedules", am)
	scheduleApi.Post("/", r.h.CreateSchedule)
	scheduleApi.Get("/rooms/:roomID", r.h.GetListScheduleByRoomID)
	scheduleApi.Get("/rooms/:roomID/tasks", r.h.GetTasksByRoomID)
	scheduleApi.Get("/tasks/:taskID", r.h.GetTaskDetail)
	scheduleApi.Get("/:scheduleID/tasks", r.h.GetListTaskByScheduleID)
	scheduleApi.Get("/:scheduleID/tasks/with-detail", r.h.GetListTaskWithDetailByScheduleID)
	scheduleApi.Get("/:scheduleID/auto/preview", r.h.AutoSchedulePreview)
	scheduleApi.Post("/:scheduleID/auto/save", r.h.AutoScheduleSave)

	taskApi := r.fr.Group("/tasks", am)
	taskApi.Post("/", r.h.PostAddTask)
	taskApi.Get("/status", r.h.GetListTaskStatus)
	taskApi.Put("/:taskID", r.h.PutEditTask)
	taskApi.Post("/dependency", r.h.CreateTaskDependency)
	taskApi.Put("/dependency", r.h.EditTaskDependency)
	taskApi.Delete("/dependency", r.h.DeleteTaskDependency)
}
