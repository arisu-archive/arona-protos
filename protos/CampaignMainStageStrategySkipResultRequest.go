package protos

type CampaignMainStageStrategySkipResultRequest struct {
	RequestPacket
	Protocol           Protocol
	PassCheckCharacter bool
	Summary            BattleSummary
}
