package protos

type CampaignPortalResponse struct {
	ResponsePacket
	Protocol                Protocol
	CampaignMainStageSaveDB CampaignMainStageSaveDB
}
