package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CharacterDB struct {
	ParcelBase
	Type                   flatdata.ParcelType
	ServerId               int64
	UniqueId               int64
	StarGrade              int32
	Level                  int32
	Exp                    int64
	FavorRank              int32
	FavorExp               int64
	PublicSkillLevel       int32
	ExSkillLevel           int32
	PassiveSkillLevel      int32
	ExtraPassiveSkillLevel int32
	LeaderSkillLevel       int32
	IsFavorite             bool
	EquipmentServerIds     []int64
	PotentialStats         map[int32]int32
}
