package friend

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
)

type FriendRoute struct {
	fr fiber.Router
	h  *handler.Handler
}

func NewFriendRoute(fr fiber.Router, h *handler.Handler) *FriendRoute {
	return &FriendRoute{
		fr, h,
	}
}

func (r *FriendRoute) InitRoute() {
	r.fr.Get("/:memberId", r.h.Friend.GetFriendsByMemberId)
}
