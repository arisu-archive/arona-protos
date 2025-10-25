package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TBGHexaObjectDB struct {
	ServerId                 int64
	UniqueId                 int64
	EncounterId              int64
	MapType                  flatdata.TBGThemaType
	Location                 HexLocation
	Activated                bool
	HitPoint                 *int32
	BeforeStoryOption        *int32
	EncounterCostAlreadyPaid bool
	IsFakeTreasure           *bool
	FixRewardUniqueIdByIndex map[int32]int64
}
