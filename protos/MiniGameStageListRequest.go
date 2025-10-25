package protos

type MiniGameStageListRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
