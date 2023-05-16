package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	"github.com/cholazzzb/amaz_corp_be/internal/app/route"
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware/auth"

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

	app := fiber.New()

	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authMiddleware := auth.CreateAuthMiddleware()

	mysqlRepo := database.NewMysqlRepository(dbMysql)
	ur := user.NewMySQLUserRepository(mysqlRepo)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)
	uRoute := route.NewUserRoute(v1, uh)
	uRoute.InitRoute(authMiddleware)

	log.Error().Err(app.Listen(":8080"))
}
