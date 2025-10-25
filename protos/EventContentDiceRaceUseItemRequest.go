package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentDiceRaceUseItemRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	DiceRaceResultType flatdata.EventContentDiceRaceResultType
}
