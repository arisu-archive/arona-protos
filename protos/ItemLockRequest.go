package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ItemLockRequest struct {
	RequestPacket
	TargetType flatdata.ParcelType `json:",omitempty,omitzero"`
	UniqueIds  []int64
}
