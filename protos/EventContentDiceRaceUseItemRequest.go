package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentDiceRaceUseItemRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	DiceRaceResultType flatdata.EventContentDiceRaceResultType `json:",omitempty,omitzero"`
}
