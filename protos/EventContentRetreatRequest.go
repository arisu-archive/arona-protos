package protos

type EventContentRetreatRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageUniqueId  int64
}
