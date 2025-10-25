package protos

type SkillLevelBatchGrowthRequestDB struct {
	SkillSlot    SkillSlot
	Level        int32
	ReplaceInfos []SelectTicketReplaceInfo
}
