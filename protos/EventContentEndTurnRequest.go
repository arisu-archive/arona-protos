package protos

type EventContentEndTurnRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageUniqueId  int64
}
