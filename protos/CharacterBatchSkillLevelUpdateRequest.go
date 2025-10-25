package protos

type CharacterBatchSkillLevelUpdateRequest struct {
	RequestPacket
	Protocol                   Protocol
	TargetCharacterDBId        int64
	SkillLevelUpdateRequestDBs []SkillLevelBatchGrowthRequestDB
}
