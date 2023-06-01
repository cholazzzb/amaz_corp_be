package heartbeat

import (
	"context"
	"time"
)

type Heartbeat struct {
	LastHeartbeat time.Time
}

type HeartbeatMap map[string]*Heartbeat

type HeartbeatRepo interface {
	GetHeartbeatMap(
		ctx context.Context,
	) (HeartbeatMap, error)
	CheckUserIdExistence(
		ctx context.Context,
		userId string,
	) (bool, error)
	UpdateToOnline(
		ctx context.Context,
		userId string,
	) error
	UpdateToOffline(
		ctx context.Context,
		userId string,
	) error
}
