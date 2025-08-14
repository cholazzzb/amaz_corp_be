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

func NewUserRouter(fr fiber.Router, h *handler.UserHandler) *UserRoute {
	return &UserRoute{
		fr, h,
	}
}

func (r *UserRoute) InitRoute(am middleware.Middleware) {
	authApi := r.fr.Group("/auth")
	authApi.Post("/v1/register", r.h.Register)

	r.fr.Group("/v1/users", am)
}
