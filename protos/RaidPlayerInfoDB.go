package protos

import (
	"time"
)

type RaidPlayerInfoDB struct {
	RaidServerId      int64
	AccountId         int64
	JoinDate          time.Time
	DamageAmount      int64
	RaidEndRewardFlag int32
	RaidPlayCount     int32
	Nickname          string
	CharacterId       int64
	CostumeId         int64
	AccountLevel      *int64
}
