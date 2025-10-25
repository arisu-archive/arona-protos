package protos

type CampaignEnterTutorialStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignTutorialStageSaveDB
}
