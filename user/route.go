package user

import "github.com/gofiber/fiber/v2"

type UserRoute struct {
	fr fiber.Router
	h  *UserHandler
}

func NewUserRoute(fr fiber.Router, h *UserHandler) *UserRoute {
	return &UserRoute{
		fr, h,
	}
}

func (r *UserRoute) InitRoute() {
	r.fr.Post("/register", r.h.Register)
	r.fr.Post("/login", r.h.Login)
}
