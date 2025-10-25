package protos

import (
	"time"
)

type ShopFreeRecruitHistoryDB struct {
	UniqueId       int64
	RecruitCount   int32
	LastUpdateDate time.Time
}
