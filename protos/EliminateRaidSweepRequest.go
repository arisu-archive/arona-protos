package protos

type EliminateRaidSweepRequest struct {
	RequestPacket
	Protocol   Protocol
	UniqueId   int64
	SweepCount int32
}
