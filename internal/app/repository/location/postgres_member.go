package location

import (
	"context"

	"github.com/google/uuid"

	locationpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location/postgresql"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
)

func (r *PostgresLocationRepository) EditMemberName(
	ctx context.Context,
	memberID string,
	memberName string,
) error {
	memberUUID, err := uuid.Parse(memberID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	_, err = r.Postgres.EditMemberName(ctx, locationpostgres.EditMemberNameParams{
		ID:   memberUUID,
		Name: memberName,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *PostgresLocationRepository) GetMemberByName(
	ctx context.Context,
	memberName string,
) (ent.MemberQuery, error) {
	result, err := r.Postgres.GetMemberByName(ctx, memberName)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.MemberQuery{}, err
	}

	return ent.MemberQuery{
		ID:     result.ID.String(),
		UserID: result.UserID.String(),
		Name:   result.Name,
		Status: result.Status,
		RoomID: result.RoomID.UUID.String(),
	}, nil
}

func (r *PostgresLocationRepository) GetMemberByID(
	ctx context.Context,
	memberID string,
) (ent.MemberQuery, error) {
	memberUUID, err := uuid.Parse(memberID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.MemberQuery{}, err
	}
	result, err := r.Postgres.GetMemberByID(ctx, memberUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.MemberQuery{}, err
	}

	return ent.MemberQuery{
		ID:     result.ID.String(),
		UserID: result.UserID.String(),
		Name:   result.Name,
		Status: result.Status,
		RoomID: result.RoomID.UUID.String(),
	}, nil
}
