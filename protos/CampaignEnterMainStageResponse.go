package protos

type CampaignEnterMainStageResponse struct {
	ResponsePacket
	Protocol   Protocol
	SaveDataDB CampaignMainStageSaveDB
}
