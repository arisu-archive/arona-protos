package protos

type EventContentStoryStageResultResponse struct {
	ResponsePacket
	Protocol                  Protocol
	CampaignStageHistoryDB    CampaignStageHistoryDB
	ParcelResultDB            ParcelResultDB
	FirstClearReward          []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
