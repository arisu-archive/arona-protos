package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MissionMultipleRewardRequest struct {
	RequestPacket
	Protocol             Protocol
	MissionCategory      flatdata.MissionCategory
	GuideMissionSeasonId *int64
	EventContentId       *int64
}
