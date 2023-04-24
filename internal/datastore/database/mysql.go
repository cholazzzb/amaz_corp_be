package database

import (
	"database/sql"
)

type MysqlRepository struct {
	Db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{
		Db: db,
	}
}
