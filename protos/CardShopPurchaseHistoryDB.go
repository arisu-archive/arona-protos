package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CardShopPurchaseHistoryDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	Rarity flatdata.Rarity `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
}
