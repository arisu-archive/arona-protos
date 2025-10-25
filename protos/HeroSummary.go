package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type HeroSummary struct {
	ServerId                  int64
	OwnerAccountId            int64
	BattleEntityId            EntityId
	CharacterId               int64
	CostumeId                 int64
	Grade                     int32
	Level                     int32
	PotentialStatLevel        map[flatdata.StatType]int32
	ExSkillLevel              int32
	PublicSkillLevel          int32
	PassiveSkillLevel         int32
	ExtraPassiveSkillLevel    int32
	FavorRank                 int32
	StatSnapshotCollection    StatSnapshotCollection
	HPRateBefore              int64
	HPRateAfter               int64
	CrowdControlCount         int32
	CrowdControlDuration      int32
	EvadeCount                int32
	DamageImmuneCount         int32
	CrowdControlImmuneCount   int32
	MaxAttackPower            int64
	AverageCriticalRate       int32
	AverageStabilityRate      int32
	AverageAccuracyRate       int32
	DeadFrame                 int32
	DamageGivenAbsorbedSum    int64
	TacticEntityType          flatdata.TacticEntityType
	GivenNumericLogs          []BattleNumericLog
	TakenNumericLogs          []BattleNumericLog
	ObstacleBattleNumericLogs []BattleNumericLog
	Equipments                []EquipmentSetting
	CharacterWeapon           *WeaponSetting
	CharacterGear             *GearSetting
	SkillCount                map[SkillSlot]int32
	KillLog                   KillLogCollection
}
