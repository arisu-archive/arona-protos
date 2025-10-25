package protos

type RaidSweepRequest struct {
	RequestPacket
	Protocol   Protocol
	UniqueId   int64
	SweepCount int64
}
