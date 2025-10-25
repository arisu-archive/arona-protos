package protos

type TimeAttackDungeonSweepRequest struct {
	RequestPacket
	Protocol   Protocol
	SweepCount int64
}
