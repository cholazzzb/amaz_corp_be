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
	buildingApi.Post("/", r.h.CreateBuilding)
	buildingApi.Get("/joined", r.h.GetBuildingsByUserID)
	buildingApi.Get("/owned", r.h.GetListMyOwnedBuilding)
	buildingApi.Get("/invitation", r.h.GetMyInvitation)
	buildingApi.Get("/all", r.h.GetBuildings) // Including public
	buildingApi.Get("/:buildingID", r.h.GetBuildingByID)
	buildingApi.Get("/:buildingId/rooms", r.h.GetRoomsByBuildingId)
	buildingApi.Get("/:buildingID/members", r.h.GetListMemberByBuildingID)

	buildingApi.Post("/invite", r.h.InviteMemberToBuilding)
	buildingApi.Post("/join", r.h.JoinBuildingById)

	buildingApi.Delete("/leave/", r.h.DeleteBuilding)

	memberApi := r.fr.Group("/members", am)
	memberApi.Put("/", r.h.EditMemberName)
	memberApi.Get("/", r.h.GetMemberByName) // (required)queryParams=name
	memberApi.Get("/:memberID", r.h.GetMemberByID)

	friendApi := r.fr.Group("/friends", am)
	friendApi.Get("/:memberId", r.h.GetFriendsByMemberId)

	roomApi := r.fr.Group("/rooms", am)
	roomApi.Post("/", r.h.CreateRoom)
	roomApi.Get("/rooms/:roomId/online", am, r.h.GetListOnlineMembers)
}
