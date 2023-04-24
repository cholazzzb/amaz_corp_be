package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app"
	"github.com/cholazzzb/amaz_corp_be/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.GetEnv()
	if config.ENV.ENVIRONMENT == "DEV" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	app.CreateApp()
}
