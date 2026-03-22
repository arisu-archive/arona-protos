package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventRewardIncreaseDB struct {
	EventTargetType flatdata.EventTargetType `json:",omitempty,omitzero"`
	Multiplier      BasisPoint
	BeginDate       MxTime `json:",omitempty,omitzero"`
	EndDate         MxTime `json:",omitempty,omitzero"`
}
