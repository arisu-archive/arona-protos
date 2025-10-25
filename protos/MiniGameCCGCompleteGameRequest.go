package protos

type MiniGameCCGCompleteGameRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
