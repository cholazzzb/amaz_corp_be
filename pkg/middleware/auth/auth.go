package auth

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/middleware"
)

type AuthHeader struct {
	Authorization string `reqHeader:"Authorization"`
}

func CreateAuthMiddleware() middleware.Middleware {
	return func(ctx *fiber.Ctx) error {
		ah := new(AuthHeader)

		if err := ctx.ReqHeaderParser(ah); err != nil {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token required")
		}

		if len(ah.Authorization) == 0 {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token required")
		}

		if !strings.Contains(ah.Authorization, "Bearer") {
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		tokenString := strings.Replace(ah.Authorization, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			} else if method != config.UserConfig.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return config.UserConfig.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token is expired")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		username, usernameOk := claims["Username"].(string)
		userID, userIDOk := claims["UserID"].(string)
		roleID, roleIDOk := claims["RoleID"].(float64)

		if !ok {
			logger.Get().Error("jwt claim failed")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !token.Valid {
			logger.Get().Error("token invalid")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !usernameOk {
			logger.Get().Error("username in token not found")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !userIDOk {
			logger.Get().Error("userID in token not found")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !roleIDOk {
			logger.Get().Error("roleID in token not found")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		ctx.Locals("Username", username)
		ctx.Locals("UserID", userID)
		ctx.Locals("RoleID", roleID)
		return ctx.Next()
	}
}

type AuthAdminHeader struct {
	Authorization string `reqHeader:"Authorization"`
	AppType       string `reqHeader:"App-Type"`
}

func CreateAuthAdminMiddleware() middleware.Middleware {
	return func(ctx *fiber.Ctx) error {

		ah := new(AuthAdminHeader)
		if err := ctx.ReqHeaderParser(ah); err != nil {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token required")
		}

		if len(ah.Authorization) == 0 {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token required")
		}

		// See roles table with id 1
		if ah.AppType != "admin" {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token required")
		}

		if !strings.Contains(ah.Authorization, "Bearer") {
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		tokenString := strings.Replace(ah.Authorization, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			} else if method != config.UserConfig.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return config.UserConfig.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).SendString("token is expired")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		username, usernameOk := claims["Username"].(string)
		userID, userIDOk := claims["UserID"].(string)
		roleID, roleIDOk := claims["RoleID"].(int32)

		if !ok {
			logger.Get().Error("jwt claim failed")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !token.Valid {
			logger.Get().Error("token invalid")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !usernameOk {
			logger.Get().Error("username in token not found")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !userIDOk {
			logger.Get().Error("userID in token not found")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		if !roleIDOk {
			logger.Get().Error("roleID in token not found")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		// See roles table with id 1
		if roleID != 1 {
			logger.Get().Error("roleID is not admin")
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}

		ctx.Locals("Username", username)
		ctx.Locals("UserID", userID)
		ctx.Locals("RoleID", roleID)
		return ctx.Next()
	}
}
