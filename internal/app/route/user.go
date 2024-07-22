package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type UserRoute struct {
	fr fiber.Router
	h  *handler.UserHandler
}

func NewUserRoute(fr fiber.Router, h *handler.UserHandler) *UserRoute {
	return &UserRoute{
		fr, h,
	}
}

func (r *UserRoute) InitRoute(am middleware.Middleware) {
	r.fr.Post("/register", r.h.Register)
	r.fr.Post("/login", r.h.Login)

	userApi := r.fr.Group("/users", am)
	userApi.Get("/", r.h.GetListUserByUsername)
	userApi.Get("/:userId/exist", r.h.CheckUserExistance)
	userApi.Get("/username/:username", r.h.GetListUserByUsername)
}
