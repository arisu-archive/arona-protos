package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type AssistCharacterDB struct {
	CharacterDB
	EchelonType             flatdata.EchelonType
	SlotNumber              int32
	AccountId               int64
	AssistRelation          AssistRelation
	AssistCharacterServerId int64
	NickName                string
	EquipmentDBs            []EquipmentDB
	WeaponDB                WeaponDB
	GearDB                  GearDB
	CostumeId               int64
	CostumeDB               CostumeDB
	IsMulligan              bool
	IsTSAInteraction        bool
	CombatStyleIndex        int32
}
