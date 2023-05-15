package repository

import (
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	friendRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/repository"
)

type Repository struct {
	Friend friendRepository.FriendRepository
}

func CreateRepository(mysqlRepo *database.MysqlRepository) *Repository {
	friendRepository := friendRepository.NewMySQLFriendRepository(mysqlRepo)

	return &Repository{
		Friend: friendRepository,
	}
}

func CreateMockRepository() *Repository {
	friendRepository := friendRepository.NewMockFriendRepository()

	return &Repository{
		Friend: friendRepository,
	}
}
