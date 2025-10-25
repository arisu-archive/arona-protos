package protos

type EventContentPortalRequest struct {
	RequestPacket
	Protocol        Protocol
	EventContentId  int64
	StageUniqueId   int64
	EchelonEntityId int64
}
