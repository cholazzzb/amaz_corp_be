package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type NotificationRoute struct {
	fr fiber.Router
}

func (r *NotificationRoute) InitRoute(am middleware.Middleware) {
	notifApi := r.fr.Group("/notifications", am)
	notifApi.Get("/list")
	notifApi.Post("/:notificationID/read")
}
