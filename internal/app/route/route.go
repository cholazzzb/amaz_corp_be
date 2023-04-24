package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	member "github.com/cholazzzb/amaz_corp_be/internal/domain/member/route"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/route"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type Route struct {
	User   *user.UserRoute
	Member *member.MemberRoute
}

func CreateRoute(api fiber.Router, h *handler.Handler, am middleware.Middleware) *Route {
	memberApi := api.Group("/member", am)

	userRoute := user.NewUserRoute(api, h)
	memberRoute := member.NewMemberRoute(memberApi, h)

	return &Route{
		User:   userRoute,
		Member: memberRoute,
	}
}

func (r *Route) InitRoute() {
	r.User.InitRoute()
	r.Member.InitRoute()
}
