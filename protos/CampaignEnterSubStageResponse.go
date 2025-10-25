package protos

type CampaignEnterSubStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignSubStageSaveDB
}
