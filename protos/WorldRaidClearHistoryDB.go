package protos

import (
	"time"
)

type WorldRaidClearHistoryDB struct {
	SeasonId          int64
	GroupId           int64
	RewardReceiveDate time.Time
}
