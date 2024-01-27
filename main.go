package main

import (
	"github.com/cholazzzb/amaz_corp_be/internal/app"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

func main() {
	config.GetEnv(".env.test")

	dbSql := app.NewSQL(app.WithMigration())
	app := app.GetApp(dbSql)

	logger.Get().Error(app.Listen(":8080").Error())
}
