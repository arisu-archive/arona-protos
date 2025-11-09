package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentShopRefreshRequest struct {
	RequestPacket
	EventContentId   int64                     `json:",omitempty,omitzero"`
	ShopCategoryType flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
