package protos

type CampaignTacticResultRequest struct {
	RequestPacket
	Protocol           Protocol
	PassCheckCharacter bool
	Summary            BattleSummary
	Hand               SkillCardHand
	SkipSummary        TacticSkipSummary
}
