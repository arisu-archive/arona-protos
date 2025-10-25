package protos

type CampaignHealRequest struct {
	RequestPacket
	Protocol              Protocol
	CampaignStageUniqueId int64
	EchelonIndex          int64
	CharacterServerId     int64
}
