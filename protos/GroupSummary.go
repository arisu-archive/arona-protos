package protos

type GroupSummary struct {
	TeamId                 int64
	LeaderEntityId         EntityId
	Heroes                 []HeroSummary
	Supporters             []HeroSummary
	UseAutoSkill           bool
	TSSInteractionServerId int64
	TSSInteractionUniqueId int64
	AssistRelations        map[int64]AssistRelation
	SkillCostSummary       SkillCostSummary
}
