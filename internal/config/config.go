package config

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	ENVIRONMENT                    string
	DB_CON_STRING                  string
	APPLICATION_NAME               string
	LOGIN_EXPIRATION_DURATION_HOUR time.Duration
	JWT_SIGNATURE_KEY              string
}

var ENV Env

func GetEnv() Env {
	env, err := godotenv.Read(".env.dev")
	if err != nil {
		log.Fatalln("failed to load env files")
	}

	ENV.ENVIRONMENT = env["ENVIRONMENT"]
	ENV.DB_CON_STRING = env["DB_CON_STRING"]
	ENV.APPLICATION_NAME = env["APPLICATION_NAME"]
	LOGIN_EXPIRATION_DURATION_HOUR, err := strconv.ParseInt(env["LOGIN_EXPIRATION_DURATION_HOUR"], 10, 64)
	if err != nil {
		log.Fatalln("failed to parse LOGIN_EXPIRATION_DURATION_HOUR from .env")
	}
	ENV.LOGIN_EXPIRATION_DURATION_HOUR = time.Duration(LOGIN_EXPIRATION_DURATION_HOUR) * time.Hour
	ENV.JWT_SIGNATURE_KEY = env["JWT_SIGNATURE_KEY"]

	CreateUserConfig(ENV)

	return ENV
}
