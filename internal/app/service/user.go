package service

import (
	"context"
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	"github.com/cholazzzb/amaz_corp_be/internal/app/model"
	repo_user "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"

	custom_logger "github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type UserService struct {
	repo   repo_user.UserRepo
	logger *slog.Logger
}

func NewUserService(
	repo repo_user.UserRepo,
) *UserService {
	sublogger := custom_logger.Get().With(slog.String("domain", "location"), slog.String("layer", "svc"))

	return &UserService{
		repo:   repo,
		logger: sublogger,
	}
}

func (s *UserService) Register(
	ctx context.Context,
	user *model.RegisterRequest,
) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.repo.Register(ctx, &model.RegisterUserCommand{
		Email:        user.Email,
		PasswordHash: string(passwordHash[:]),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Timezone:     user.Timezone,
	})

	return err
}
