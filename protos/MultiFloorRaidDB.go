package protos

import (
	"time"
)

type MultiFloorRaidDB struct {
	SeasonId          int64
	ClearedDifficulty int32
	LastClearDate     time.Time
	RewardDifficulty  int32
	LastRewardDate    time.Time
	ClearBattleFrame  int32
}
