package protos

import (
	"time"
)

type MissionHistoryDB struct {
	MissionUniqueId int64
	CompleteTime    time.Time
	Expired         bool
}
