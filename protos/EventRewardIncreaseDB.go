package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"time"
)

type EventRewardIncreaseDB struct {
	EventTargetType flatdata.EventTargetType `json:",omitempty,omitzero"`
	Multiplier BasisPoint `json:",omitempty,omitzero"`
	BeginDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
}
