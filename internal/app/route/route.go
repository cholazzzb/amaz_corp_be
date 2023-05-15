package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	friend "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/route"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/route"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type Route struct {
	User   *user.UserRoute
	Friend *friend.FriendRoute
}

func CreateRoute(api fiber.Router, h *handler.Handler, am middleware.Middleware) *Route {

	friendApi := api.Group("/friends", am)
	friendRoute := friend.NewFriendRoute(friendApi, h)

	return &Route{
		Friend: friendRoute,
	}
}

func (r *Route) InitRoute() {
	r.Friend.InitRoute()
}
