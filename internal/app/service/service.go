package service

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	friend "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/service"
	member "github.com/cholazzzb/amaz_corp_be/internal/domain/member/service"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/service"
)

type Service struct {
	User   *user.UserService
	Member *member.MemberService
	Friend *friend.FriendService
}

func CreateService(repo *repository.Repository) *Service {
	userService := user.NewUserService(repo)
	memberService := member.NewMemberService(repo)
	friendService := friend.NewFriendService(repo)

	return &Service{
		User:   userService,
		Member: memberService,
		Friend: friendService,
	}
}
