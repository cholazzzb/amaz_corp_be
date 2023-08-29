package migrator

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

func newMigrator() *migrate.FileMigrationSource {
	var dir string
	switch config.ENV.DB_TYPE {
	case "postgres":
		dir = "./migration/postgresql"
	case "mysql":
		dir = "./migration/mysql"
	default:
		logger.Get().Error("config.ENV.DB_TYPE is not recognized")
		panic("config.ENV.DB_TYPE is not recognized")
	}

	return &migrate.FileMigrationSource{
		Dir: dir,
	}
}

func MigrateUp(dbSql *sql.DB) {
	n, err := migrate.Exec(dbSql, config.ENV.DB_TYPE, newMigrator(), migrate.Up)
	if err != nil {
		logger.Get().Error("failed to migrate database")
		panic("failed to migrate database")
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
