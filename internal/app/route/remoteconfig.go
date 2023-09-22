package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type RemoteConfigRoute struct {
	fr fiber.Router
	h  *handler.RemoteConfigHandler
}

func NewRemoteConfigRoute(
	fr fiber.Router,
	h *handler.RemoteConfigHandler,
) *RemoteConfigRoute {
	return &RemoteConfigRoute{
		fr, h,
	}
}

func (r *RemoteConfigRoute) InitRoute(am middleware.Middleware) {
	remoteconfigApi := r.fr.Group("/remoteconfig")
	remoteconfigApi.Get("/apk-version", r.h.GetAPKVersion)
}
