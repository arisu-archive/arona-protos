package protos

type MiniGameEnterStageRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	UniqueId       int64
}
