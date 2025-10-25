package protos

type CampaignRestartMainStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignMainStageSaveDB
}
