package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BillingTransactionStartByYostarResponse struct {
	ResponsePacket
	Protocol                  Protocol
	PurchaseCount             int64
	PurchaseResetDate         time.Time
	PurchaseOrderId           int64
	MXSeedKey                 string
	PurchaseServerTag         flatdata.PurchaseServerTag
	PurchaseServerCallbackUrl string
}
