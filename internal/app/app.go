package app

import (
	"database/sql"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	hbRepo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/heartbeat"
	locRepo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location"
	rcRepo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/remoteconfig"
	schRepo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	"github.com/cholazzzb/amaz_corp_be/internal/app/route"
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/heartbeat"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware/auth"
	"github.com/cholazzzb/amaz_corp_be/pkg/migrator"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var app *fiber.App

func GetApp(dbSql *sql.DB) *fiber.App {
	if app == nil {
		lock.Lock()
		defer lock.Unlock()

		if app == nil {
			// TODO: Write the log into file

			opt, err := redis.ParseURL(config.ENV.REDIS_CON_STRING)
			if err != nil {
				logger.Get().Error(err.Error())
				panic("failed to connect redis database")
			}
			rds := redis.NewClient(opt)
			defer rds.Close()

			app = fiber.New()

			app.Use(fiberLogger.New(fiberLogger.Config{
				TimeFormat: "2006-01-02T15:04:05-0700",
			}))

			app.Use(cors.New())

			api := app.Group("/api")
			v1 := api.Group("/v1")

			authMiddleware := auth.CreateAuthMiddleware()
			authAdminMiddleware := auth.CreateAuthAdminMiddleware()

			sqlRepo := database.NewSqlRepository(dbSql)
			redisRepo := database.NewRedisRepository(rds)

			ur := user.NewPostgresUserRepository(sqlRepo)
			us := service.NewUserService(ur)
			uh := handler.NewUserHandler(us)
			uRoute := route.NewUserRoute(v1, uh)
			uRoute.InitRoute(authMiddleware, authAdminMiddleware)

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

			sr := schRepo.NewPostgresScheduleRepository(sqlRepo)
			scr := schRepo.NewRedisScheduleRepository(redisRepo)
			ss := service.NewScheduleService(sr, scr)
			sh := handler.NewScheduleHandler(ss)
			sRoute := route.NewScheduleRoute(v1, sh)
			sRoute.InitRoute(authMiddleware)

			rcr := rcRepo.NewPostgresRemoteConfigRepository(sqlRepo)
			rcs := service.NewRemoteConfigService(rcr)
			rch := handler.NewRemoteConfigHandler(rcs)
			rcRoute := route.NewRemoteConfigRoute(v1, rch)
			rcRoute.InitRoute(authMiddleware)
		}
	}

	return app
}

func NewSQL(options ...func(*sql.DB)) *sql.DB {
	dbSql, err := sql.Open(config.ENV.DB_TYPE, config.ENV.DB_CON_STRING)
	if err != nil {
		logger.Get().Error(err.Error())
		panic("failed to connect sql database")
	}

	for _, opt := range options {
		opt(dbSql)
	}

	return dbSql
}

func WithMigration() func(*sql.DB) {
	return func(dbSql *sql.DB) {
		migrator.MigrateUp(dbSql)
	}
}
