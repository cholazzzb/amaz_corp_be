package heartbeat

import (
	"context"
	"time"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/heartbeat"
)

type HeartBeatScheduler struct {
	repo repo.HeartbeatRepo
}

func NewHeartBeatScheduler(repo repo.HeartbeatRepo) *HeartBeatScheduler {
	return &HeartBeatScheduler{repo}
}

func (sch *HeartBeatScheduler) Schedule(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			hbm, _ := sch.repo.GetHeartbeatMap(context.Background())
			for memberId, heartbeat := range hbm {
				momentsAgo := time.Now().Add(-time.Duration(interval))
				isBefore := heartbeat.LastHeartbeat.Before(momentsAgo)

				if isBefore {
					sch.repo.UpdateToOffline(context.Background(), memberId)
				}

			}
		}
	}
}
