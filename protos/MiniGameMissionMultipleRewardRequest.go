package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MiniGameMissionMultipleRewardRequest struct {
	RequestPacket
	MissionCategory flatdata.MissionCategory `json:",omitempty,omitzero"`
	EventContentId  int64                    `json:",omitempty,omitzero"`
}
