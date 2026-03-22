package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type TBGBoardSaveDB struct {
	AccountId                      int64                 `json:",omitempty,omitzero"`
	EventContentId                 int64                 `json:",omitempty,omitzero"`
	Round                          int32                 `json:",omitempty,omitzero"`
	ThemaIndex                     int32                 `json:",omitempty,omitzero"`
	CurrentThemaMapType            flatdata.TBGThemaType `json:",omitempty,omitzero"`
	MainMap                        TBGHexaMapDB
	HiddenMap                      TBGHexaMapDB
	Player                         TBGPlayerDB
	Encounter                      TBGEncounterDB
	BestClearRecord                *mapx.OrderedMap[int32, TBGThemaClearRecord]
	HiddenTreasureRecord           []int32
	HiddenPotalOpenConditionRecord []int32
}
