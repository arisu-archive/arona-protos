package protos

type EventContentTacticResultRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	PassCheckCharacter bool
	Summary            BattleSummary
	Hand               SkillCardHand
	SkipSummary        TacticSkipSummary
}
