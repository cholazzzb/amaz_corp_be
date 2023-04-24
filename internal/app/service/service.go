package service

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	member "github.com/cholazzzb/amaz_corp_be/internal/domain/member/service"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/service"
)

type Service struct {
	User   *user.UserService
	Member *member.MemberService
}

func CreateService(repo *repository.Repository) *Service {
	userService := user.NewUserService(repo)
	memberService := member.NewMemberService(repo)

	return &Service{
		User:   userService,
		Member: memberService,
	}
}
