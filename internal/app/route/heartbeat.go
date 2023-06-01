package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type HeartbeatRoute struct {
	fr fiber.Router
	h  *handler.HeartbeatHandler
}

func NewHeartbeatRoute(fr fiber.Router, h *handler.HeartbeatHandler) *HeartbeatRoute {
	return &HeartbeatRoute{
		fr, h,
	}
}

func (r *HeartbeatRoute) InitRoute(am middleware.Middleware) {
	heartbeatApi := r.fr.Group("/heartbeats", am)
	heartbeatApi.Get("/pulse", r.h.Pulse)
	heartbeatApi.Get("/check/:userId", r.h.GetStatusByUserId)
}
