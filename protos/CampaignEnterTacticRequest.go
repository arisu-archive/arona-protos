package protos

type CampaignEnterTacticRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
	EchelonIndex  int64
	EnemyIndex    int64
}
