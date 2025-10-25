package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type SelectTicketReplaceInfo struct {
	MaterialType flatdata.ParcelType
	MaterialId   int64
	TicketItemId int64
	Amount       int32
}
