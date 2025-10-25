package protos

import (
	"time"
)

type PurchaseCountDB struct {
	ShopCashId      int64
	PurchaseCount   int32
	ResetDate       time.Time
	PurchaseDate    *time.Time
	ManualResetDate *time.Time
}
