package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanAssistSlotDB struct {
	EchelonType      flatdata.EchelonType
	SlotNumber       int64
	CharacterDBId    int64
	DeployDate       time.Time
	TotalRentCount   int64
	CombatStyleIndex int32
}
