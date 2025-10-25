package protos

type EventContentEnterSubStageResponse struct {
	ResponsePacket
	Protocol               Protocol
	ParcelResultDB         ParcelResultDB
	SaveDataDB             EventContentSubStageSaveDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
