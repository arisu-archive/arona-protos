package protos

type CampaignConfirmTutorialStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
