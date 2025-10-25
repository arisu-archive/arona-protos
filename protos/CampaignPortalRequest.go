package protos

type CampaignPortalRequest struct {
	RequestPacket
	Protocol        Protocol
	StageUniqueId   int64
	EchelonEntityId int64
}
