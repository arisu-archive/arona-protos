package protos

type MiniGameShootingSweepRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	UniqueId       int64
	SweepCount     int64
}
