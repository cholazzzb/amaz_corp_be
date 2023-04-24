package user

import (
	"context"
	"errors"
	"os/exec"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Username string
}

type UserService struct {
	repo   *repository.Repository
	logger zerolog.Logger
}

type Token string

func NewUserService(
	repo *repository.Repository,
) *UserService {
	sublogger := log.With().Str("layer", "service").Str("package", "user").Logger()

	return &UserService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *UserService) RegisterUser(ctx context.Context, username, password string) error {
	salt, err := exec.Command("uuidgen").Output()
	if err != nil {
		svc.logger.Error().Err(err)
		return errors.New("failed to generate uuid")
	}
	saltedPassword := append(
		[]byte(password),
		salt...,
	)

	hashedPassword, err := bcrypt.GenerateFromPassword(
		saltedPassword,
		bcrypt.DefaultCost,
	)
	if err != nil {
		svc.logger.Error().Err(err)
		return errors.New("failed to hash password")
	}
	newUserParams := mysql.CreateUserParams{
		Username: username,
		Password: string(hashedPassword[:]),
		Salt:     string(salt[:]),
	}

	if err := svc.repo.User.CreateUser(ctx, newUserParams); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (svc *UserService) authenticateUser(ctx context.Context, username, password string) (bool, error) {
	result, err := svc.repo.User.GetUser(ctx, username)
	if err != nil {
		svc.logger.Error().Err(err).Msg("username not found")
		return false, errors.New("username not found")
	}

	saltedPassword := append(
		[]byte(password),
		[]byte(result.Salt)...,
	)

	err = bcrypt.CompareHashAndPassword(
		[]byte(result.Password),
		saltedPassword,
	)

	if err != nil {
		return false, errors.New("wrong Password")
	}
	return true, nil
}

func (svc *UserService) Login(ctx context.Context, username, password string) (Token, error) {
	isAuthentic, err := svc.authenticateUser(ctx, username, password)
	if !isAuthentic || err != nil {
		return "", errors.New("failed to authenticate user")
	}

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.UserConfig.APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.UserConfig.LOGIN_EXPIRATION_DURATION)),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.UserConfig.JWT_SIGNATURE_KEY)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return Token(signedToken), nil
}
