package protos

type CampaignEndTurnRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
