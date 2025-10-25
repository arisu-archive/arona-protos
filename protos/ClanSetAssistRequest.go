package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanSetAssistRequest struct {
	RequestPacket
	Protocol         Protocol
	EchelonType      flatdata.EchelonType
	SlotNumber       int32
	CharacterDBId    int64
	CombatStyleIndex int32
}
