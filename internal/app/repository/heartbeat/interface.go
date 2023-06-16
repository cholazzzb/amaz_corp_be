package heartbeat

import (
	"context"
	"time"
)

type Heartbeat struct {
	LastHeartbeat time.Time
}

// map[userId]*Heartbeat
type HeartbeatMap map[int64]*Heartbeat

type HeartbeatRepo interface {
	GetHeartbeatMap(
		ctx context.Context,
	) (HeartbeatMap, error)
	CheckUserIdExistence(
		ctx context.Context,
		userId int64,
	) (bool, error)
	UpdateToOnline(
		ctx context.Context,
		userId int64,
	) error
	UpdateToOffline(
		ctx context.Context,
		userId int64,
	) error
}
