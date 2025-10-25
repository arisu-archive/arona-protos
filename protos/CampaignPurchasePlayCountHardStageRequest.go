package protos

type CampaignPurchasePlayCountHardStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
