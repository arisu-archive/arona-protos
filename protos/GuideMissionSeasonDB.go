package protos

import (
	"time"
)

type GuideMissionSeasonDB struct {
	SeasonId                  int64
	LoginCount                int64
	StartDate                 time.Time
	LoginDate                 time.Time
	IsComplete                bool
	IsFinalMissionComplete    bool
	CollectionItemReceiveDate *time.Time
}
