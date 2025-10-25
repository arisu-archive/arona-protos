package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ShopRefreshRequest struct {
	RequestPacket
	Protocol         Protocol
	ShopCategoryType flatdata.ShopCategoryType
}
