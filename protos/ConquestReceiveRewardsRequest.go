package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestReceiveRewardsRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	Difficulty flatdata.StageDifficulty `json:",omitempty,omitzero"`
	Step int32 `json:",omitempty,omitzero"`
}
