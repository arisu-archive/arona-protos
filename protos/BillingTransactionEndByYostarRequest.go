package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BillingTransactionEndByYostarRequest struct {
	RequestPacket
	PurchaseOrderId int64 `json:",omitempty,omitzero"`
	EndType flatdata.BillingTransactionEndType `json:",omitempty,omitzero"`
}
