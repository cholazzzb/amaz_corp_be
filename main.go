package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/database"
	"github.com/cholazzzb/amaz_corp_be/internal/member"
	"github.com/cholazzzb/amaz_corp_be/internal/user"
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

	mysqlRepo := database.NewMysqlRepository(dbMysql)

	app := fiber.New()

	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: "02-Jan-2006",
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	userRepository := user.NewUserRepository(mysqlRepo)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	userRoute := user.NewUserRoute(v1, userHandler)
	userRoute.InitRoute()

	authMiddleware := auth.CreateAuthMiddleware()

	memberRepository := member.NewMySQLMemberRepository(mysqlRepo)
	memberService := member.NewMemberService(memberRepository)
	memberHandler := member.NewMemberHandler(memberService)

	memberApi := v1.Group("/member", authMiddleware)
	memberRoute := member.NewMemberRoute(memberApi, memberHandler)
	memberRoute.InitRoute()

	log.Error().Err(app.Listen(":8080"))
}
