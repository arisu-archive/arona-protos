package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestReceiveRewardsRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Difficulty     flatdata.StageDifficulty
	Step           int32
}
