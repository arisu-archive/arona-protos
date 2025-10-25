package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MiniGameMissionMultipleRewardRequest struct {
	RequestPacket
	Protocol        Protocol
	MissionCategory flatdata.MissionCategory
	EventContentId  int64
}
