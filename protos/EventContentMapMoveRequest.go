package protos

type EventContentMapMoveRequest struct {
	RequestPacket
	Protocol        Protocol
	EventContentId  int64
	StageUniqueId   int64
	EchelonEntityId int64
	DestPosition    HexLocation
}
