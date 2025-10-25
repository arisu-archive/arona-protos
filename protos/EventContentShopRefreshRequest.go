package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentShopRefreshRequest struct {
	RequestPacket
	Protocol         Protocol
	EventContentId   int64
	ShopCategoryType flatdata.ShopCategoryType
}
