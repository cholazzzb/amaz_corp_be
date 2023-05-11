package member

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
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
	r.fr.Post("", r.h.Member.CreateMemberByUsername)
}
