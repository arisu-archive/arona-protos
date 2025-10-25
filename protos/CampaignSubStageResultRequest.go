package protos

type CampaignSubStageResultRequest struct {
	RequestPacket
	Protocol           Protocol
	PassCheckCharacter bool
	Summary            BattleSummary
}
