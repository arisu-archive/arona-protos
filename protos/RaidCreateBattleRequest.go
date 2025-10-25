package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidCreateBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	RaidUniqueId  int64
	IsPractice    bool
	Tags          []int32
	IsPublic      bool
	Difficulty    flatdata.Difficulty
	AssistUseInfo ClanAssistUseInfo
}
