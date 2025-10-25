package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClearDeckCharacterDB struct {
	UniqueId         int64
	StarGrade        int32
	Level            int32
	SlotNumber       int32
	HasWeapon        bool
	SquadType        flatdata.SquadType
	WeaponStarGrade  int32
	CombatStyleIndex int32
}
