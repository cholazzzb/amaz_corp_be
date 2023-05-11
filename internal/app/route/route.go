package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	friend "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/route"
	member "github.com/cholazzzb/amaz_corp_be/internal/domain/member/route"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/route"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type Route struct {
	User   *user.UserRoute
	Member *member.MemberRoute
	Friend *friend.FriendRoute
}

func CreateRoute(api fiber.Router, h *handler.Handler, am middleware.Middleware) *Route {
	userRoute := user.NewUserRoute(api, h)

	memberApi := api.Group("/members", am)
	memberRoute := member.NewMemberRoute(memberApi, h)

	friendApi := api.Group("/friends", am)
	friendRoute := friend.NewFriendRoute(friendApi, h)

	return &Route{
		User:   userRoute,
		Member: memberRoute,
		Friend: friendRoute,
	}
}

func (r *Route) InitRoute() {
	r.User.InitRoute()
	r.Member.InitRoute()
	r.Friend.InitRoute()
}
