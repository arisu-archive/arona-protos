package protos

type TimeAttackDungeonCharacterDB struct {
	ServerId         int64
	UniqueId         int64
	CostumeId        int64
	StarGrade        int32
	Level            int32
	HasWeapon        bool
	WeaponDB         WeaponDB
	IsAssist         bool
	CombatStyleIndex int32
}
