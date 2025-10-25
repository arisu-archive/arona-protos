package protos

type CampaignConfirmMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
