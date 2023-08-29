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
	scheduleApi.Get("/rooms/:roomID", r.h.GetScheduleIDByRoomID)
	scheduleApi.Get("/tasks/:taskID", r.h.GetTaskDetail)
	scheduleApi.Get("/:scheduleID/tasks", r.h.GetListTaskByScheduleID)

	taskApi := r.fr.Group("/tasks", am)
	taskApi.Post("/", r.h.PostAddTask)
	taskApi.Put("/:taskID", r.h.PutEditTask)
}
