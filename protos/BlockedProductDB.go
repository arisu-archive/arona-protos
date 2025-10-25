package protos

import (
	"time"
)

type BlockedProductDB struct {
	CashProductId   int64
	MarketBlockType ShopCashBlockType
	BeginDate       time.Time
	EndDate         time.Time
}
