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
	buildingApi.Get("/", r.h.GetBuildingsByUserID)
	buildingApi.Get("/all", r.h.GetBuildings)
	buildingApi.Get("/:buildingId/rooms", r.h.GetRoomsByBuildingId)
	buildingApi.Get("/:buildingID/members", r.h.GetListMemberByBuildingID)

	buildingApi.Post("/join", r.h.JoinBuildingById)

	buildingApi.Delete("/leave/", r.h.DeleteBuilding)

	memberApi := r.fr.Group("/members", am)
	memberApi.Get("/:name", r.h.GetMemberByName)

	friendApi := r.fr.Group("/friends", am)
	friendApi.Get("/:memberId", r.h.GetFriendsByMemberId)

	r.fr.Get("/rooms/:roomId/online", am, r.h.GetListOnlineMembers)
}
