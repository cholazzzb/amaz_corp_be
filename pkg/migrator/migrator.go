package migrator

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	db *sql.DB
}

func newMigrator() *migrate.FileMigrationSource {
	return &migrate.FileMigrationSource{
		Dir: "./migration",
	}
}

func MigrateUp(dbMysql *sql.DB) {
	n, err := migrate.Exec(dbMysql, "mysql", newMigrator(), migrate.Up)
	if err != nil {
		log.Panic().Err(err).Msg("failed to migrate mysql database")
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
