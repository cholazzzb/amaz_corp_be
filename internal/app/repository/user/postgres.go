package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/google/uuid"

	userpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type PostgresUserRepository struct {
	Postgres *userpostgres.Queries
	logger   *slog.Logger
}

func NewPostgresUserRepository(sqlRepo *database.SqlRepository) *PostgresUserRepository {
	sublogger := logger.Get().With("domain", "user", "layer", "repo")

	queries := userpostgres.New(sqlRepo.Db)
	return &PostgresUserRepository{Postgres: queries, logger: sublogger}
}

func (r *PostgresUserRepository) GetUser(
	ctx context.Context,
	params string,
) (user.User, error) {
	result, err := r.Postgres.GetUser(ctx, params)
	if err != nil {
		r.logger.Error(err.Error())
		return user.User{}, err
	}
	return user.User{
		ID:       result.ID.String(),
		Username: result.Username,
		Password: result.Password,
		Salt:     result.Salt,
	}, nil
}

func (r *PostgresUserRepository) GetListUserByUsername(
	ctx context.Context,
	username string,
) ([]user.UserQuery, error) {
	out := []user.UserQuery{}

	likeParams := fmt.Sprint("%", username, "%")
	res, err := r.Postgres.GetListUserByUsername(ctx, likeParams)
	if err != nil {
		r.logger.Error(err.Error())
		return out, err
	}

	for _, usr := range res {
		out = append(out, user.UserQuery{
			ID:       usr.ID.String(),
			Username: usr.Username,
		})
	}

	return out, nil
}

func (r *PostgresUserRepository) GetUserExistance(
	ctx context.Context,
	username string,
) (bool, error) {
	exist, err := r.Postgres.GetUserExistance(ctx, username)
	if err != nil {
		r.logger.Error(err.Error())
		return true, err
	}
	return exist, nil
}

func (r *PostgresUserRepository) CreateUser(
	ctx context.Context,
	params user.UserCommand,
) error {
	_, err := r.Postgres.CreateUser(ctx, userpostgres.CreateUserParams{
		Username:  params.Username,
		Password:  params.Password,
		Salt:      params.Salt,
		ProductID: params.ProductID,
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) GetProductByUserID(
	ctx context.Context,
	userID string,
) (user.ProductQuery, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err.Error())
		return user.ProductQuery{}, err
	}

	product, err := r.Postgres.GetProductByUserID(ctx, userUUID)
	if err != nil {
		return user.ProductQuery{}, err
	}
	return user.ProductQuery{
		ID:   product.ID,
		Name: product.Name,
	}, nil
}

func (r *PostgresUserRepository) GetListProduct(
	ctx context.Context,
) ([]user.ProductQuery, error) {
	out := []user.ProductQuery{}
	products, err := r.Postgres.GetListProduct(ctx)
	if err != nil {
		return []user.ProductQuery{}, err
	}

	for _, prd := range products {
		out = append(out, user.ProductQuery{
			ID:   prd.ID,
			Name: prd.Name,
		})
	}
	return out, nil
}

func (r *PostgresUserRepository) GetListFeatureByProductID(
	ctx context.Context,
	productID int32,
) ([]user.FeatureQuery, error) {
	out := []user.FeatureQuery{}

	feats, err := r.Postgres.GetListFeatureByProductID(ctx, productID)
	if err != nil {
		return []user.FeatureQuery{}, err
	}

	for _, feat := range feats {
		out = append(out, user.FeatureQuery{
			ID:   feat.ID.String(),
			Name: feat.Name,
		})
	}
	return out, nil
}
