package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"time"
)

type MonthlyProductPurchaseDB struct {
	ProductId int64 `json:",omitempty,omitzero"`
	PurchaseDate time.Time `json:",omitempty,omitzero"`
	LastDailyRewardDate *time.Time `json:",omitempty,omitzero"`
	RewardEndDate *time.Time `json:",omitempty,omitzero"`
	ProductTagType flatdata.ProductTagType `json:",omitempty,omitzero"`
}
