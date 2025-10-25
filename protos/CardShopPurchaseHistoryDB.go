package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CardShopPurchaseHistoryDB struct {
	EventContentId int64
	Rarity         flatdata.Rarity
	PurchaseCount  int64
}
