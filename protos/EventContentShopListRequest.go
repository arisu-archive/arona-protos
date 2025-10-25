package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentShopListRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	CategoryList   []flatdata.ShopCategoryType
}
