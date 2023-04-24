package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	"github.com/cholazzzb/amaz_corp_be/internal/app/route"
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware/auth"

	_ "github.com/go-sql-driver/mysql"
)

func CreateApp() {
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
	repository := repository.CreateRepository(mysqlRepo)
	service := service.CreateService(repository)
	handler := handler.CreateHandler(service)
	route := route.CreateRoute(v1, handler, authMiddleware)
	route.InitRoute()

	log.Error().Err(app.Listen(":8080"))
}
