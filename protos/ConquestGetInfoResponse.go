package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestGetInfoResponse struct {
	ResponsePacket
	Protocol                 Protocol
	ConquestInfoDB           ConquestInfoDB
	ConquestedTileDBs        []ConquestTileDB
	ConquestObjectDBsWrapper []ConquestEventObjectDB
	ConquestEchelonDBs       []ConquestEchelonDB
	DifficultyToStepDict     map[flatdata.StageDifficulty]int32
	IsFirstEnter             bool
	DisplayInfos             []ConquestDisplayInfo
}
