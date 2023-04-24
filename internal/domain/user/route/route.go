package user

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	fr fiber.Router
	h  *handler.Handler
}

func NewUserRoute(fr fiber.Router, h *handler.Handler) *UserRoute {
	return &UserRoute{
		fr, h,
	}
}

func (r *UserRoute) InitRoute() {
	r.fr.Post("/register", r.h.User.Register)
	r.fr.Post("/login", r.h.User.Login)
}
