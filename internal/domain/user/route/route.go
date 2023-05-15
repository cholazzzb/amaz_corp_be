package user

import (
	"github.com/gofiber/fiber/v2"

	hdl "github.com/cholazzzb/amaz_corp_be/internal/domain/user/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type UserRoute struct {
	fr fiber.Router
	h  *hdl.UserHandler
}

func NewUserRoute(fr fiber.Router, h *hdl.UserHandler) *UserRoute {
	return &UserRoute{
		fr, h,
	}
}

func (r *UserRoute) InitRoute(am middleware.Middleware) {
	r.fr.Post("/register", r.h.Register)
	r.fr.Post("/login", r.h.Login)

	memberApi := r.fr.Group("/members", am)
	memberApi.Get("/:name", r.h.GetMemberByName)
	memberApi.Post("", r.h.CreateMemberByUsername)
}
