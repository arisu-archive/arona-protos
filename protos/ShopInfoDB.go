package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ShopInfoDB struct {
	EventContentId      int64
	Category            flatdata.ShopCategoryType
	ManualRefreshCount  *int64
	IsRefresh           bool
	ShopGroupType       flatdata.ShopGroupType
	NextAutoRefreshDate *time.Time
	LastAutoRefreshDate *time.Time
	ShopProductList     []ShopProductDB
}
