package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TacticalRelayRankingInfoRequest struct {
	RequestPacket
	SeasonId  int64                           `json:",omitempty,omitzero"`
	StageType flatdata.TacticalRelayStageType `json:",omitempty,omitzero"`
}
