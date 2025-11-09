package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BattlePassMissionMultipleRewardRequest struct {
	RequestPacket
	MissionCategory flatdata.MissionCategory `json:",omitempty,omitzero"`
	BattlePassId    int64                    `json:",omitempty,omitzero"`
}
