package protos

import (
	"time"
)

type PickupFirstGetHistoryDB struct {
	ShopRecruitId int64
	LogDate       time.Time
}
