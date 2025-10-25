package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestConquerDeployEchelonRequest struct {
	RequestPacket
	Protocol          Protocol
	EventContentId    int64
	Difficulty        flatdata.StageDifficulty
	TileUniqueId      int64
	EchelonDB         EchelonDB
	ClanAssistUseInfo ClanAssistUseInfo
}
