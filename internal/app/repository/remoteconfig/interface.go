package remoteconfig

import "context"

type RemoteConfigRepo interface {
	RemoteConfigCommand
	RemoteConfigQuery
}

type RemoteConfigCommand interface {
	CreateRemoteConfig(
		ctx context.Context,
		key string,
		value string,
	) error

	UpdateRemoteConfig(
		ctx context.Context,
		key string,
		value string,
	) error
}

type RemoteConfigQuery interface {
	GetRemoteConfigByKey(
		ctx context.Context,
		key string,
	) (string, error)
}
