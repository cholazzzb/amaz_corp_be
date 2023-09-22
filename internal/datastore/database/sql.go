package database

import (
	"database/sql"
)

type SqlRepository struct {
	Db *sql.DB
}

func NewSqlRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{
		Db: db,
	}
}
