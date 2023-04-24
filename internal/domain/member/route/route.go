package member

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/gofiber/fiber/v2"
)

type MemberRoute struct {
	fr fiber.Router
	h  *handler.Handler
}

func NewMemberRoute(fr fiber.Router, h *handler.Handler) *MemberRoute {
	return &MemberRoute{
		fr, h,
	}
}

func (r *MemberRoute) InitRoute() {
	r.fr.Get("/:name", r.h.Member.GetMemberByName)
}
