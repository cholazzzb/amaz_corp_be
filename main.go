package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	userHandler "github.com/cholazzzb/amaz_corp_be/internal/domain/user/handler"
	userRepo "github.com/cholazzzb/amaz_corp_be/internal/domain/user/repository"
	userRoute "github.com/cholazzzb/amaz_corp_be/internal/domain/user/route"
	userService "github.com/cholazzzb/amaz_corp_be/internal/domain/user/service"
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
	ur := userRepo.NewMySQLUserRepository(mysqlRepo)
	us := userService.NewUserService(ur)
	uh := userHandler.NewUserHandler(us)
	uRoute := userRoute.NewUserRoute(v1, uh)
	uRoute.InitRoute(authMiddleware)

	log.Error().Err(app.Listen(":8080"))
}
