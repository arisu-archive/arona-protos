package protos

type CampaignDeployEchelonResponse struct {
	ResponsePacket
	Protocol   Protocol
	SaveDataDB CampaignMainStageSaveDB
}
