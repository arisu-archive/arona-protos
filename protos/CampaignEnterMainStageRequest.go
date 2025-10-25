package protos

type CampaignEnterMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
