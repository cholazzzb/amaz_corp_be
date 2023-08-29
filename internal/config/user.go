package config

import (
	"strconv"
	"time"

	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type userConfig struct {
	APPLICATION_NAME          string
	LOGIN_EXPIRATION_DURATION time.Duration
	JWT_SIGNING_METHOD        *jwt.SigningMethodHMAC
	JWT_SIGNATURE_KEY         []byte
}

var UserConfig userConfig

func CreateUserConfig(env map[string]string) {

	APPLICATION_NAME, ok := env["APPLICATION_NAME"]
	if !ok {
		logger.Get().Error("failed to parse APPLICATION_NAME from .env")
		panic("failed to parse APPLICATION_NAME from .env")
	}

	LOGIN_EXPIRATION_DURATION_HOUR, err := strconv.ParseInt(env["LOGIN_EXPIRATION_DURATION_HOUR"], 10, 64)
	if err != nil {
		logger.Get().Error("failed to parse LOGIN_EXPIRATION_DURATION_HOUR from .env")
		panic("failed to parse LOGIN_EXPIRATION_DURATION_HOUR from .env")
	}
	LOGIN_EXPIRATION_DURATION := time.Duration(LOGIN_EXPIRATION_DURATION_HOUR) * time.Hour

	JWT_SIGNATURE_KEY := env["JWT_SIGNATURE_KEY"]

	UserConfig = userConfig{
		APPLICATION_NAME:          APPLICATION_NAME,
		LOGIN_EXPIRATION_DURATION: LOGIN_EXPIRATION_DURATION,
		JWT_SIGNING_METHOD:        jwt.SigningMethodHS256,
		JWT_SIGNATURE_KEY:         []byte(JWT_SIGNATURE_KEY),
	}
}
