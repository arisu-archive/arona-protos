package protos

type CampaignRetreatRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
