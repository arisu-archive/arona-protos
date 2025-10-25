package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanAssistUseInfo struct {
	CharacterAccountId int64
	CharacterDBId      int64
	EchelonType        flatdata.EchelonType
	SlotNumber         int32
	AssistRelation     AssistRelation
	EchelonSlotType    int32
	EchelonSlotIndex   int32
	CombatStyleIndex   int32
	IsMulligan         bool
	IsTSAInteraction   bool
}
