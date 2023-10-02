package service

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Username string
	UserId   string
}

type UserService struct {
	repo   repo.UserRepo
	logger *slog.Logger
}

type Token string

func NewUserService(
	repo repo.UserRepo,
) *UserService {
	sublogger := logger.Get().With(slog.String("domain", "user"), slog.String("layer", "svc"))

	return &UserService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *UserService) RegisterUser(ctx context.Context, username, password string) error {
	exist, err := svc.CheckUserExistance(ctx, username)

	if err != nil {
		svc.logger.Error(err.Error())
		return errors.New("failed to check username existence")
	}

	if exist {
		return errors.New("username already taken")
	}

	uuidSalt, err := uuid.NewUUID()
	if err != nil {
		svc.logger.Error(err.Error())
		return errors.New("failed to generate uuid")
	}

	salt := uuidSalt.String()
	saltedPassword := append(
		[]byte(password),
		salt...,
	)

	hashedPassword, err := bcrypt.GenerateFromPassword(
		saltedPassword,
		bcrypt.DefaultCost,
	)
	if err != nil {
		svc.logger.Error(err.Error())
		return errors.New("failed to hash password")
	}
	newUserParams := user.UserCommand{
		Username:  username,
		Password:  string(hashedPassword[:]),
		Salt:      string(salt[:]),
		ProductID: 1, // TODO: Find better way to write this, for now: 1 = is free, the default value
	}

	if err := svc.repo.CreateUser(ctx, newUserParams); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (svc *UserService) authenticateUser(ctx context.Context, username, password string) (bool, string, error) {
	result, err := svc.repo.GetUser(ctx, username)
	if err != nil {
		svc.logger.Error(err.Error())
		return false, "", errors.New("username not found")
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
		svc.logger.Error(err.Error())
		return false, "", errors.New("wrong Password")
	}
	return true, result.ID, nil
}

func (svc *UserService) Login(ctx context.Context, username, password string) (Token, error) {
	isAuthentic, userId, err := svc.authenticateUser(ctx, username, password)
	if !isAuthentic || err != nil {
		svc.logger.Error(err.Error())
		return "", errors.New("failed to authenticate user")
	}

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.UserConfig.APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.UserConfig.LOGIN_EXPIRATION_DURATION)),
		},
		Username: username,
		UserId:   userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.UserConfig.JWT_SIGNATURE_KEY)
	if err != nil {
		svc.logger.Error(err.Error())
		return "", errors.New("failed to sign token")
	}

	return Token(signedToken), nil
}

func (svc *UserService) CheckUserExistance(ctx context.Context, username string) (bool, error) {
	exist, err := svc.repo.GetUserExistance(ctx, username)
	if err != nil {
		svc.logger.Error(err.Error())
		return false, err
	}
	return exist, nil
}

func (svc *UserService) GetProductByUserID(
	ctx context.Context,
	userID string,
) (user.ProductQuery, error) {
	out, err := svc.repo.GetProductByUserID(ctx, userID)
	if err != nil {
		return user.ProductQuery{}, err
	}

	return out, nil
}

func (svc *UserService) GetListProduct(
	ctx context.Context,
) ([]user.ProductQuery, error) {
	out, err := svc.repo.GetListProduct(ctx)
	if err != nil {
		return []user.ProductQuery{}, err
	}

	return out, nil
}

func (svc *UserService) GetListFeatureByProductID(
	ctx context.Context,
	productID int32,
) ([]user.FeatureQuery, error) {
	out, err := svc.repo.GetListFeatureByProductID(ctx, productID)
	if err != nil {
		return []user.FeatureQuery{}, err
	}

	return out, nil
}
