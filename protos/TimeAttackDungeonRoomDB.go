package protos

import (
	"time"
)

type TimeAttackDungeonRoomDB struct {
	AccountId         int64
	SeasonId          int64
	RoomId            int64
	CreateDate        time.Time
	RewardDate        time.Time
	IsPractice        bool
	SweepHistoryDates []time.Time
	BattleHistoryDBs  []TimeAttackDungeonBattleHistoryDB
}
