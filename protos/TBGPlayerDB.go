package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TBGPlayerDB struct {
	Location             HexLocation
	EventContentId       int64
	HitPoint             int32
	DiceId               int64
	DiceProbModifyParams map[flatdata.TBGProbModifyCondition]int32
	Items                []TBGItemDB
	TemporaryItem        TBGItemDB
	ItemEffects          []TBGItemEffectDB
}
