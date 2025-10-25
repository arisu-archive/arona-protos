package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ShopProductDB struct {
	EventContentId     int64
	ShopExcelId        int64
	Category           flatdata.ShopCategoryType
	DisplayOrder       int64
	PurchaseCount      int64
	PurchaseCountLimit int64
	Price              int64
	ProductType        ShopProductType
}
