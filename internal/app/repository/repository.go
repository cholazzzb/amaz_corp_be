package repository

import (
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	memberRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/member/repository"
	userRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/user/repository"
)

type Repository struct {
	User   userRepository.UserRepository
	Member memberRepository.MemberRepository
}

func CreateRepository(mysqlRepo *database.MysqlRepository) *Repository {
	userRepository := userRepository.NewMySQLUserRepository(mysqlRepo)
	memberRepository := memberRepository.NewMySQLMemberRepository(mysqlRepo)

	return &Repository{
		User:   userRepository,
		Member: memberRepository,
	}
}

func CreateMockRepository() *Repository {
	userRepository := userRepository.NewMockUserRepository()
	memberRepository := memberRepository.NewMockMemberRepository()

	return &Repository{
		User:   userRepository,
		Member: memberRepository,
	}
}
