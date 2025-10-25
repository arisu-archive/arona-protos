package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestMainStoryGetInfoResponse struct {
	ResponsePacket
	Protocol             Protocol
	ConquestInfoDB       ConquestInfoDB
	ConquestedTileDBs    []ConquestTileDB
	DifficultyToStepDict map[flatdata.StageDifficulty]int32
	IsFirstEnter         bool
}
