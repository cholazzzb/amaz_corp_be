package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	ENVIRONMENT                    string
	DB_TYPE                        string
	DB_CON_STRING                  string
	REDIS_CON_STRING               string
	LOGIN_EXPIRATION_DURATION_HOUR time.Duration
}

var ENV Env

func GetEnv(envLocation string) Env {
	env, err := godotenv.Read(envLocation)
	if err != nil {
		log.Fatalln("failed to load env files", err)
	}

	ENV.ENVIRONMENT = env["ENVIRONMENT"]
	ENV.DB_TYPE = env["DB_TYPE"]
	ENV.DB_CON_STRING = env["DB_CON_STRING"]
	ENV.REDIS_CON_STRING = env["REDIS_CON_STRING"]

	CreateUserConfig(env)
	CreateHeartbeatConfig(env)

	return ENV
}
