package service

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	friend "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/service"
)

type Service struct {
	Friend *friend.FriendService
}

func CreateService(repo *repository.Repository) *Service {
	friendService := friend.NewFriendService(repo)

	return &Service{
		Friend: friendService,
	}
}
