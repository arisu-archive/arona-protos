package protos

type CampaignTutorialStageResultResponse struct {
	ResponsePacket
	Protocol               Protocol
	CampaignStageHistoryDB CampaignStageHistoryDB
	ParcelResultDB         ParcelResultDB
	ClearReward            []ParcelInfo
	FirstClearReward       []ParcelInfo
}
