package config

import (
	"log"
	"strconv"
	"time"
)

type heartbeatConfig struct {
	CHECK_INTERVAL time.Duration
}

var Heartbeat heartbeatConfig

func CreateHeartbeatConfig(env map[string]string) {
	HEARTBEAT_CHECK_INTERVAL, err := strconv.ParseInt(env["HEARTBEAT_CHECK_INTERVAL"], 10, 64)
	if err != nil {
		log.Fatalln("failed to parse HEARTBEAT_CHECK_INTERVAL from .env")
	}
	hbInterval := time.Duration(HEARTBEAT_CHECK_INTERVAL) * time.Second

	Heartbeat = heartbeatConfig{CHECK_INTERVAL: hbInterval}
}
