package protos

type MiniGameCCGSweepRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	SweepCount     int32
}
