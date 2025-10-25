package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TBGItemEffectDB struct {
	ITBGItemEffectDB
	ItemUniqueId           int64
	ItemType               flatdata.TBGItemType
	EffectType             flatdata.TBGItemEffectType
	RemainEncounterCounter int32
}
