package protos

import (
	"time"
)

type MissionProgressDB struct {
	MissionUniqueId    int64
	Complete           bool
	StartTime          time.Time
	ProgressParameters map[int64]int64
}
