package protos

type EventContentEnterMainGroundStageResponse struct {
	ResponsePacket
	Protocol               Protocol
	ParcelResultDB         ParcelResultDB
	SaveDataDB             EventContentMainGroundStageSaveDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
