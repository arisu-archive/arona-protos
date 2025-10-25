package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TBGBoardSaveDB struct {
	AccountId                      int64
	EventContentId                 int64
	Round                          int32
	ThemaIndex                     int32
	CurrentThemaMapType            flatdata.TBGThemaType
	MainMap                        TBGHexaMapDB
	HiddenMap                      TBGHexaMapDB
	Player                         TBGPlayerDB
	Encounter                      TBGEncounterDB
	BestClearRecord                map[int32]TBGThemaClearRecord
	HiddenTreasureRecord           []int32
	HiddenPotalOpenConditionRecord []int32
}
