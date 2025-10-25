package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventRewardIncreaseDB struct {
	EventTargetType flatdata.EventTargetType
	Multiplier      BasisPoint
	BeginDate       time.Time
	EndDate         time.Time
}
