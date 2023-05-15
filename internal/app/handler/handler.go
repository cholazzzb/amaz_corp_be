package handler

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	friend "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/handler"
)

type Handler struct {
	Friend *friend.FriendHandler
}

func CreateHandler(service *service.Service) *Handler {
	friendHandler := friend.NewFriendHandler(service)

	return &Handler{
		Friend: friendHandler,
	}
}
