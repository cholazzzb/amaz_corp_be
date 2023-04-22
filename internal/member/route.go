package member

import "github.com/gofiber/fiber/v2"

type MemberRoute struct {
	fr fiber.Router
	h  *MemberHandler
}

func NewMemberRoute(fr fiber.Router, h *MemberHandler) *MemberRoute {
	return &MemberRoute{
		fr, h,
	}
}

func (r *MemberRoute) InitRoute() {
	r.fr.Get("/:name", r.h.GetMemberByName)
}
