package route

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

type ReportRoute struct {
	fr fiber.Router
	h  *handler.ReportHandler
}

func NewReportRoute(fr fiber.Router, h *handler.ReportHandler) *ReportRoute {
	return &ReportRoute{
		fr, h,
	}
}

func (r *ReportRoute) InitRoute(am middleware.Middleware) {
	reportApi := r.fr.Group("/reports", am)
	reportApi.Get("/schedules/:scheduleID", r.h.GetReportBySchedule)

	reportApi.Get("/rooms/:roomID", r.h.GetReportBySchedule)

	reportApi.Get("/members/:memberID", r.h.GetReportBySchedule)

}
