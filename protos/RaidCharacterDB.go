package protos

type RaidCharacterDB struct {
	ServerId         int64
	UniqueId         int64
	StarGrade        int32
	Level            int32
	SlotIndex        int32
	AccountId        int64
	IsAssist         bool
	HasWeapon        bool
	WeaponStarGrade  int32
	CostumeId        int64
	CombatStyleIndex int32
}
