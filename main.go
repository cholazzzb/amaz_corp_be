package main

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app"
	"github.com/cholazzzb/amaz_corp_be/internal/config"

	custom_logger "github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

func main() {
	config.GetEnv(".env.dev")

	dbSql := app.NewSQL(app.WithMigration())
	app := app.GetApp(dbSql)

	custom_logger.Get().Error(app.Listen(":8080").Error())
}
