package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type PurchaseOrderDB struct {
	ShopCashId      int64
	StatusCode      flatdata.PurchaseStatusCode
	PurchaseOrderId int64
}
