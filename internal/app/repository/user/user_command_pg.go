package repo_user

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/app/model"
	repository_user "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/postgresql"
	"github.com/cholazzzb/amaz_corp_be/pkg/mapper"
	"github.com/google/uuid"
)

func (r *UserRepository) Register(
	ctx context.Context,
	user *model.RegisterUserCommand,
) error {
	_, err := r.Pg.RegisterNewUser(ctx, repository_user.RegisterNewUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FirstName:    mapper.ToNullString(user.FirstName),
		LastName:     mapper.ToNullString(user.LastName),
		Timezone:     mapper.ToNullString(user.Timezone),
	})
	return err
}

func (r *UserRepository) Update(
	ctx context.Context,
	id string,
	user *model.UserProfileCommand,
) (*model.UserProfileCommandRes, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		r.logger.Error("failed to parse uuid", err)
		return nil, err
	}

	result, err := r.Pg.UpdateUserProfile(ctx, repository_user.UpdateUserProfileParams{
		ID:        uuid,
		FirstName: mapper.ToNullString(user.FirstName),
		LastName:  mapper.ToNullString(user.LastName),
		Timezone:  mapper.ToNullString(user.Timezone),
	})

	if err != nil {
		return nil, err
	}

	return &model.UserProfileCommandRes{
		Id:        result.ID.String(),
		FirstName: mapper.ToString(result.FirstName),
		LastName:  mapper.ToString(result.LastName),
		Timezone:  mapper.ToString(result.Timezone),
		UpdatedAt: mapper.ToTime(result.UpdatedAt).String(),
	}, nil
}

func (r *UserRepository) Deactivate(
	ctx context.Context,
	id string,
) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		r.logger.Error("failed to parse uuid", err)
		return err
	}

	_, err = r.Pg.DeactivateUserAccount(ctx, uuid)

	return err
}
