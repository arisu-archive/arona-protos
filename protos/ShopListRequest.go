package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ShopListRequest struct {
	RequestPacket
	Protocol     Protocol
	CategoryList []flatdata.ShopCategoryType
}
