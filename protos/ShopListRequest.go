package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ShopListRequest struct {
	RequestPacket
	CategoryList []flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
