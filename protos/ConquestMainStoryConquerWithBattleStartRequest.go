package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestMainStoryConquerWithBattleStartRequest struct {
	RequestPacket
	Protocol          Protocol
	EventContentId    int64
	Difficulty        flatdata.StageDifficulty
	TileUniqueId      int64
	EchelonNumber     *int64
	ClanAssistUseInfo ClanAssistUseInfo
}
