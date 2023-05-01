package main

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app"
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.GetEnv()
	if config.ENV.ENVIRONMENT == "DEV" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	dbMysql, err := sql.Open("mysql", config.ENV.DB_CON_STRING)
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect mysql database")
	}

	mysqlRepo := database.NewMysqlRepository(dbMysql)
	repository := repository.CreateRepository(mysqlRepo)

	app.CreateApp(repository)
}
