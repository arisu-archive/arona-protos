package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestConquerRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Difficulty     flatdata.StageDifficulty
	TileUniqueId   int64
	TileRewardId   int64
}
