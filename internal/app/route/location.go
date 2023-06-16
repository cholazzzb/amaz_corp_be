package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type LocationRoute struct {
	fr fiber.Router
	h  *handler.LocationHandler
}

func NewLocationRoute(fr fiber.Router, h *handler.LocationHandler) *LocationRoute {
	return &LocationRoute{
		fr, h,
	}
}

func (r *LocationRoute) InitRoute(am middleware.Middleware) {
	buildingApi := r.fr.Group("/buildings", am)
	buildingApi.Get("/", r.h.GetBuildingsByMemberId)
	buildingApi.Get("/all", r.h.GetBuildings)
	buildingApi.Get("/:buildingId/rooms", r.h.GetRoomsByBuildingId)

	buildingApi.Post("/join", r.h.JoinRoomById)

	r.fr.Get("/rooms/:roomId/online", am, r.h.GetListOnlineMembers)
}
