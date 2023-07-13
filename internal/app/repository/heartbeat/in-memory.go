package heartbeat

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type InMemoryHeartbeatRepo struct {
	heartbeatMap HeartbeatMap
	mu           *sync.Mutex
}

func NewInMemoryHeartbeatRepo() *InMemoryHeartbeatRepo {
	return &InMemoryHeartbeatRepo{
		heartbeatMap: HeartbeatMap{},
		mu:           &sync.Mutex{},
	}
}

func (r *InMemoryHeartbeatRepo) GetHeartbeatMap(
	ctx context.Context,
) (HeartbeatMap, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.heartbeatMap, nil
}

func (r *InMemoryHeartbeatRepo) CheckUserIdExistence(
	ctx context.Context,
	userId string,
) (bool, error) {
	hbmap, err := r.GetHeartbeatMap(ctx)

	if err != nil {
		return false, err
	}
	_, ok := hbmap[userId]

	return ok, nil
}

func (r *InMemoryHeartbeatRepo) UpdateToOnline(
	ctx context.Context,
	userId string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.heartbeatMap[userId] = &Heartbeat{
		LastHeartbeat: time.Now(),
	}

	return nil
}

func (r *InMemoryHeartbeatRepo) UpdateToOffline(
	ctx context.Context,
	userId string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.heartbeatMap[userId]

	if !ok {
		return fmt.Errorf("error when update to offline, can't find memberId: %v", userId)
	}

	delete(r.heartbeatMap, userId)

	return nil
}
