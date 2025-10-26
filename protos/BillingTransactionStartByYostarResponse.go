package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"time"
)

type BillingTransactionStartByYostarResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
	PurchaseResetDate time.Time `json:",omitempty,omitzero"`
	PurchaseOrderId int64 `json:",omitempty,omitzero"`
	MXSeedKey string `json:",omitempty,omitzero"`
	PurchaseServerTag flatdata.PurchaseServerTag `json:",omitempty,omitzero"`
	PurchaseServerCallbackUrl string `json:",omitempty,omitzero"`
}
