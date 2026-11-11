package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type TBGBoardSaveDB struct {
	AccountId                      int64                                         `json:"aid,omitempty,omitzero"`
	EventContentId                 int64                                         `json:"ecid,omitempty,omitzero"`
	Round                          int32                                         `json:"rnd,omitempty,omitzero"`
	ThemaIndex                     int32                                         `json:"idx,omitempty,omitzero"`
	CurrentThemaMapType            flatdata.TBGThemaType                         `json:"map_type,omitempty,omitzero"`
	MainMap                        *TBGHexaMapDB                                 `json:"map_main,omitempty,omitzero"`
	HiddenMap                      *TBGHexaMapDB                                 `json:"map_hidden,omitempty,omitzero"`
	Player                         *TBGPlayerDB                                  `json:"usr,omitempty,omitzero"`
	Encounter                      TBGEncounterDBPoly                            `json:"enc"`
	BestClearRecord                *mapx.OrderedMap[int32, *TBGThemaClearRecord] `json:"clearlog"`
	HiddenTreasureRecord           []int32                                       `json:"htr_idx"`
	HiddenPotalOpenConditionRecord []int32                                       `json:"hcr_idx"`
}
