package protos

type ScenarioTacticResultRequest struct {
	RequestPacket
	Protocol           Protocol
	PassCheckCharacter bool
	Summary            BattleSummary
	Hand               SkillCardHand
	SkipSummary        TacticSkipSummary
}
