package app

import (
	"database/sql"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	hbRepo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/heartbeat"
	locRepo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location"
	"github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	"github.com/cholazzzb/amaz_corp_be/internal/app/route"
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/heartbeat"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware/auth"
	"github.com/cholazzzb/amaz_corp_be/pkg/migrator"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var app *fiber.App

func GetApp() *fiber.App {
	if app == nil {
		lock.Lock()
		defer lock.Unlock()

		if app == nil {
			if config.ENV.ENVIRONMENT == "DEV" {
				log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			}

			dbSql, err := sql.Open(config.ENV.DB_TYPE, config.ENV.DB_CON_STRING)
			if err != nil {
				log.Panic().Err(err).Msg("failed to connect sql database")
			}

			migrator.MigrateUp(dbSql)

			opt, err := redis.ParseURL(config.ENV.REDIS_CON_STRING)
			if err != nil {
				log.Panic().Err(err).Msg("failed to connect redis database")
			}
			redis.NewClient(opt)

			app = fiber.New()

			app.Use(fiberLogger.New(fiberLogger.Config{
				TimeFormat: "2006-01-02T15:04:05-0700",
			}))

			api := app.Group("/api")
			v1 := api.Group("/v1")

			authMiddleware := auth.CreateAuthMiddleware()

			sqlRepo := database.NewSqlRepository(dbSql)
			ur := user.NewPostgresUserRepository(sqlRepo)
			us := service.NewUserService(ur)
			uh := handler.NewUserHandler(us)
			uRoute := route.NewUserRoute(v1, uh)
			uRoute.InitRoute(authMiddleware)

			hbr := hbRepo.NewInMemoryHeartbeatRepo()
			go heartbeat.NewHeartBeatScheduler(hbr).Schedule(
				config.Heartbeat.CHECK_INTERVAL,
			)
			hrs := service.NewHeartbeatService(hbr)
			hrh := handler.NewHeartBeatHandler(hrs)
			hrRoute := route.NewHeartbeatRoute(v1, hrh)
			hrRoute.InitRoute(authMiddleware)

			lr := locRepo.NewPostgresLocationRepository(sqlRepo)
			ls := service.NewLocationService(hrs, us, lr)
			lh := handler.NewLocationHandler(ls)
			lRoute := route.NewLocationRoute(v1, lh)
			lRoute.InitRoute(authMiddleware)
		}
	}

	return app
}
