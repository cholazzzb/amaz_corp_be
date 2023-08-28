package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Username string
	UserId   string
}

type UserService struct {
	repo   repo.UserRepo
	logger zerolog.Logger
}

type Token string

func NewUserService(
	repo repo.UserRepo,
) *UserService {
	sublogger := log.With().Str("layer", "service").Str("package", "user").Logger()

	return &UserService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *UserService) RegisterUser(ctx context.Context, username, password string) error {
	exist, err := svc.CheckUserExistance(ctx, username)

	if err != nil {
		svc.logger.Error().Err(err)
		return errors.New("failed to check username existence")
	}

	if exist {
		return errors.New("username already taken")
	}

	uuidSalt, err := uuid.NewV7()
	if err != nil {
		svc.logger.Error().Err(err)
		return errors.New("failed to generate uuid")
	}

	userId, err := uuid.NewV7()
	if err != nil {
		svc.logger.Error().Err(err)
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
		svc.logger.Error().Err(err)
		return errors.New("failed to hash password")
	}
	newUserParams := user.User{
		ID:       userId.String(),
		Username: username,
		Password: string(hashedPassword[:]),
		Salt:     string(salt[:]),
	}

	if err := svc.repo.CreateUser(ctx, newUserParams); err != nil {
		svc.logger.Error().Err(err)
		return errors.New("failed to create user")
	}

	return nil
}

func (svc *UserService) authenticateUser(ctx context.Context, username, password string) (bool, string, error) {
	result, err := svc.repo.GetUser(ctx, username)
	if err != nil {
		svc.logger.Error().Err(err).Msg("username not found")
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
		return false, "", errors.New("wrong Password")
	}
	return true, result.ID, nil
}

func (svc *UserService) Login(ctx context.Context, username, password string) (Token, error) {
	isAuthentic, userId, err := svc.authenticateUser(ctx, username, password)
	if !isAuthentic || err != nil {
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
		return "", errors.New("failed to sign token")
	}

	return Token(signedToken), nil
}

func (svc *UserService) CheckUserExistance(ctx context.Context, username string) (bool, error) {
	exist, err := svc.repo.GetUserExistance(ctx, username)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (svc *UserService) GetMemberByName(ctx context.Context, name string) (user.Member, error) {
	member, err := svc.repo.GetMemberByName(ctx, name)
	if err != nil {
		return member, fmt.Errorf("cannot find member with name %s", name)
	}
	return member, nil
}

// TODO: Use transaction here
func (svc *UserService) CreateMember(ctx context.Context, memberName string, username string) (user.Member, error) {
	userData, err := svc.repo.GetUser(ctx, username)
	if err != nil {
		svc.logger.Error().Err(err)
		return user.Member{}, fmt.Errorf("cannot found user with username %s", username)
	}

	memberId, err := uuid.NewV7()
	if err != nil {
		svc.logger.Error().Err(err)
		return user.Member{}, errors.New("failed to generate memberId")
	}

	newMember, err := svc.repo.CreateMember(ctx, user.Member{
		ID:     memberId.String(),
		Name:   memberName,
		Status: "new member",
	}, userData.ID)

	if err != nil {
		svc.logger.Error().Err(err)
		return user.Member{}, fmt.Errorf("failed to create member %v", newMember)
	}
	return newMember, nil
}

func (svc *UserService) GetFriendsByMemberId(ctx context.Context, userId string) ([]user.Member, error) {
	fs, err := svc.repo.GetFriendsByUserId(ctx, userId)
	if err != nil {
		svc.logger.Error().Err(err)
		return nil, fmt.Errorf("cannot find friends with name %s", fs)
	}
	return fs, nil
}
