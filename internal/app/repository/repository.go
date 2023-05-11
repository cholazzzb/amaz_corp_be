package repository

import (
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	friendRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/repository"
	memberRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/member/repository"
	userRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/user/repository"
)

type Repository struct {
	User   userRepository.UserRepository
	Member memberRepository.MemberRepository
	Friend friendRepository.FriendRepository
}

func CreateRepository(mysqlRepo *database.MysqlRepository) *Repository {
	userRepository := userRepository.NewMySQLUserRepository(mysqlRepo)
	memberRepository := memberRepository.NewMySQLMemberRepository(mysqlRepo)
	friendRepository := friendRepository.NewMySQLFriendRepository(mysqlRepo)

	return &Repository{
		User:   userRepository,
		Member: memberRepository,
		Friend: friendRepository,
	}
}

func CreateMockRepository() *Repository {
	userRepository := userRepository.NewMockUserRepository()
	memberRepository := memberRepository.NewMockMemberRepository()
	friendRepository := friendRepository.NewMockFriendRepository()

	return &Repository{
		User:   userRepository,
		Member: memberRepository,
		Friend: friendRepository,
	}
}
