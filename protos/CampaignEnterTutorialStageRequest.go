package protos

type CampaignEnterTutorialStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
