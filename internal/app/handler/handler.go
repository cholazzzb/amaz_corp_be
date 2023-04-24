package handler

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	member "github.com/cholazzzb/amaz_corp_be/internal/domain/member/handler"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/handler"
)

type Handler struct {
	User   *user.UserHandler
	Member *member.MemberHandler
}

func CreateHandler(service *service.Service) *Handler {
	userHandler := user.NewUserHandler(service)
	memberHandler := member.NewMemberHandler(service)

	return &Handler{
		User:   userHandler,
		Member: memberHandler,
	}
}
