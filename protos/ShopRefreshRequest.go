package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ShopRefreshRequest struct {
	RequestPacket
	ShopCategoryType flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
