package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BillingTransactionStartByYostarResponse struct {
	ResponsePacket
	PurchaseCount             int64                      `json:",omitempty,omitzero"`
	PurchaseResetDate         MxTime                     `json:",omitempty,omitzero"`
	PurchaseOrderId           int64                      `json:",omitempty,omitzero"`
	MXSeedKey                 string                     `json:",omitempty,omitzero"`
	PurchaseServerTag         flatdata.PurchaseServerTag `json:",omitempty,omitzero"`
	PurchaseServerCallbackUrl string                     `json:",omitempty,omitzero"`
}
