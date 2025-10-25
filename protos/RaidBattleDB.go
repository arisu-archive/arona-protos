package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidBattleDB struct {
	ContentType        flatdata.ContentType
	RaidUniqueId       int64
	RaidBossIndex      int32
	CurrentBossHP      int64
	CurrentBossGroggy  int64
	CurrentBossAIPhase int64
	BIEchelon          string
	IsClear            bool
	RaidMembers        RaidMemberCollection
	SubPartsHPs        []int64
}
