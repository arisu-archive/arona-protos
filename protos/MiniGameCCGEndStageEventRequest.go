package protos

type MiniGameCCGEndStageEventRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
