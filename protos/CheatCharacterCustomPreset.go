package protos

type CheatCharacterCustomPreset struct {
	UniqueId            int64
	StarGrade           int32
	Level               int32
	ExSkillLevel        int32
	PublicSkillLevel    int32
	PassiveSkillLevel   int32
	ExPassiveSkillLevel int32
	Equipments          []CheatEquipmentCustomPreset
	Weapon              CheatWeaponCustomPreset
}
