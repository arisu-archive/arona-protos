package protos

type MiniGameCCGEndStageDualRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Summary        MiniGameCCGSummary
}
