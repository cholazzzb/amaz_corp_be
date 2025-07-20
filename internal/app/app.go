package app

import (
	"database/sql"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"

	"github.com/cholazzzb/amaz_corp_be/internal/app/handler"
	repo_user "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	"github.com/cholazzzb/amaz_corp_be/internal/app/route"
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	database "github.com/cholazzzb/amaz_corp_be/internal/datastore"

	custom_logger "github.com/cholazzzb/amaz_corp_be/pkg/logger"
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
				custom_logger.Get().Error(err.Error())
				panic("failed to connect redis database")
			}
			rds := redis.NewClient(opt)
			defer rds.Close()

			app = fiber.New()

			app.Use(logger.New(logger.Config{
				TimeFormat: "2006-01-02T15:04:05-0700",
			}))

			app.Use(cors.New())

			api := app.Group("/api")

			authMiddleware := auth.CreateAuthMiddleware()

			sqlRepo := database.NewSqlRepository(dbSql)

			// User
			ur := repo_user.NewRepository(sqlRepo)
			us := service.NewUserService(ur)
			uh := handler.NewUserHandler(us)
			uRoute := route.NewUserRouter(api, uh)
			uRoute.InitRoute(authMiddleware)
		}
	}

	return app
}

func NewSQL(options ...func(*sql.DB)) *sql.DB {
	dbSql, err := sql.Open(config.ENV.DB_TYPE, config.ENV.DB_CON_STRING)
	if err != nil {
		custom_logger.Get().Error(err.Error())
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
