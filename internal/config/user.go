package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userConfig struct {
	APPLICATION_NAME          string
	LOGIN_EXPIRATION_DURATION time.Duration
	JWT_SIGNING_METHOD        *jwt.SigningMethodHMAC
	JWT_SIGNATURE_KEY         []byte
}

var UserConfig userConfig

func CreateUserConfig(env Env) {
	UserConfig = userConfig{
		APPLICATION_NAME:          env.APPLICATION_NAME,
		LOGIN_EXPIRATION_DURATION: env.LOGIN_EXPIRATION_DURATION_HOUR,
		JWT_SIGNING_METHOD:        jwt.SigningMethodHS256,
		JWT_SIGNATURE_KEY:         []byte(env.JWT_SIGNATURE_KEY),
	}
}
