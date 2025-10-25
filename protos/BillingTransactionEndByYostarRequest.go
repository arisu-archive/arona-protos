package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BillingTransactionEndByYostarRequest struct {
	RequestPacket
	Protocol        Protocol
	PurchaseOrderId int64
	EndType         flatdata.BillingTransactionEndType
}
