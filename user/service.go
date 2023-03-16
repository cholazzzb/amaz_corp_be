package user

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"time"

	"github.com/cholazzzb/amaz_corp_be/config"
	mysql "github.com/cholazzzb/amaz_corp_be/user/mysql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	APPLICATION_NAME          = config.ENV.APPLICATION_NAME
	LOGIN_EXPIRATION_DURATION = config.ENV.LOGIN_EXPIRATION_DURATION_HOUR
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte(config.ENV.JWT_SIGNATURE_KEY)
)

type UserClaims struct {
	jwt.RegisteredClaims
	Username string
}

type UserService struct {
	repo *UserRepository
}

type Token string

func NewUserService(
	repo *UserRepository,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) RegisterUser(ctx context.Context, username, password string) error {
	salt, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return errors.New("failed to hash password")
	}
	newUserParams := mysql.CreateUserParams{
		Username: username,
		Password: string(hashedPassword[:]),
		Salt:     string(salt[:]),
	}

	if err := svc.repo.CreateUser(ctx, newUserParams); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (svc *UserService) authenticateUser(ctx context.Context, username, password string) (bool, error) {
	result, err := svc.repo.GetUser(ctx, username)
	if err != nil {
		log.Println("svc/user/authenticateUser: username not found")
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
			Issuer:    APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return Token(signedToken), nil
}
