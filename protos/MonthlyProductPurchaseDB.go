package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MonthlyProductPurchaseDB struct {
	ProductId           int64
	PurchaseDate        time.Time
	LastDailyRewardDate *time.Time
	RewardEndDate       *time.Time
	ProductTagType      flatdata.ProductTagType
}
