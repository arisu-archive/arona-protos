package protos

import (
	"time"
)

type RaidDetailDB struct {
	RaidUniqueId int64
	EndDate      time.Time
	DamageTable  []RaidPlayerInfoDB
}
