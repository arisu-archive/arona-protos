package protos

import (
	"time"
)

type ArenaBattleDB struct {
	ArenaBattleServerId int64
	Season              int64
	Group               int64
	BattleStartTime     time.Time
	BattleEndTime       time.Time
	Seed                int64
	AttackingUserDB     ArenaUserDB
	DefendingUserDB     ArenaUserDB
	BattleSummary       BattleSummary
}
