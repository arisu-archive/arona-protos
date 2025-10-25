package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BattlePassMissionMultipleRewardRequest struct {
	RequestPacket
	Protocol        Protocol
	MissionCategory flatdata.MissionCategory
	BattlePassId    int64
}
