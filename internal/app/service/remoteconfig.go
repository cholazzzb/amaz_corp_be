package service

import (
	"context"
	"fmt"
	"log/slog"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/remoteconfig"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/remoteconfig"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type RemoteConfigService struct {
	repo   repo.RemoteConfigRepo
	logger *slog.Logger
}

func NewRemoteConfigService(repo repo.RemoteConfigRepo) *RemoteConfigService {
	sublogger := logger.Get().With(
		slog.String("domain", "remoteconfig"),
		slog.String("layer", "svc"),
	)
	return &RemoteConfigService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *RemoteConfigService) CreateRemoteConfig(
	ctx context.Context,
	key,
	value string,
) error {
	return svc.CreateRemoteConfig(ctx, key, value)
}

func (svc *RemoteConfigService) UpdateRemoteConfig(
	ctx context.Context,
	key,
	value string,
) error {
	return svc.UpdateRemoteConfig(ctx, key, value)
}

func (svc *RemoteConfigService) getRemoteConfigByKey(
	ctx context.Context,
	key string,
) (ent.RemoteConfigQuery, error) {
	// TODO: use cache using redis
	val, err := svc.repo.GetRemoteConfigByKey(ctx, key)
	if err != nil {
		return ent.RemoteConfigQuery{}, fmt.Errorf("repo, remoteconfig. err:%w", err)
	}

	return ent.RemoteConfigQuery{
		Key:   key,
		Value: val,
	}, nil
}

func (svc *RemoteConfigService) GetAPKVersion(
	ctx context.Context,
) (ent.APKVersionQuery, error) {
	rc, err := svc.getRemoteConfigByKey(ctx, "apk-version")
	if err != nil {
		return ent.APKVersionQuery{}, fmt.Errorf("apk-version. err:%w", err)
	}

	return ent.APKVersionQuery{
		ApkVersion: rc.Value,
	}, nil
}
