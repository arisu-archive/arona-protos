package protos

type CampaignRestartMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
