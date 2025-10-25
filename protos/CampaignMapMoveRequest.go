package protos

type CampaignMapMoveRequest struct {
	RequestPacket
	Protocol        Protocol
	StageUniqueId   int64
	EchelonEntityId int64
	DestPosition    HexLocation
}
