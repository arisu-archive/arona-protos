package protos

type MiniGameCCGEnterStageRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	NodeId         int64
}
