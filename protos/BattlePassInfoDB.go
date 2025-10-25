package protos

import (
	"time"
)

type BattlePassInfoDB struct {
	BattlePassId                      int64
	PassLevel                         int32
	PassExp                           int64
	PurchaseGroupId                   int64
	ReceiveRewardLevel                int32
	ReceivePurchaseRewardLevel        int32
	WeeklyPassExp                     int64
	LastWeeklyPassExpLimitRefreshDate time.Time
}
