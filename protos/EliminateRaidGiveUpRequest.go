package protos

type EliminateRaidGiveUpRequest struct {
	RequestPacket
	Protocol     Protocol
	RaidServerId int64
	IsPractice   bool
}
