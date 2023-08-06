package main

import (
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
)

func main() {
	config.GetEnv(".env")

	app := app.GetApp()

	log.Error().Err(app.Listen(":8080"))
}
