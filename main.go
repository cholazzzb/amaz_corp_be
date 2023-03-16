package main

import (
	"database/sql"
	"log"

	"github.com/cholazzzb/amaz_corp_be/config"
	"github.com/cholazzzb/amaz_corp_be/database"
	"github.com/cholazzzb/amaz_corp_be/user"
	"github.com/gofiber/fiber/v2"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.GetEnv()

	dbMysql, err := sql.Open("mysql", config.ENV.DB_CON_STRING)
	if err != nil {
		log.Panicln("failed to connect mysql database", err)
	}

	mysqlRepo := database.NewMysqlRepository(dbMysql)

	app := fiber.New()
	api := app.Group("/api")

	v1 := api.Group("/v1")

	userRepository := user.NewUserRepository(mysqlRepo)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	userRoute := user.NewUserRoute(v1, userHandler)
	userRoute.InitRoute()

	log.Fatal(app.Listen(":8080"))
}
