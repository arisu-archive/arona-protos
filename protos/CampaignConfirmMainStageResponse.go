package protos

type CampaignConfirmMainStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignMainStageSaveDB
}
