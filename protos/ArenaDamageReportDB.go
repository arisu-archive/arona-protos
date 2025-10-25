package protos

import (
	"time"
)

type ArenaDamageReportDB struct {
	ArenaBattleServerId   int64
	WinnerAccountServerId int64
	AttackerUserDB        ArenaUserDB
	DefenderUserDB        ArenaUserDB
	BattleEndTime         time.Time
	AttackerDamageReport  map[int64]int64
	DefenderDamageReport  map[int64]int64
}
