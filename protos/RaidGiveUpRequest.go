package protos

type RaidGiveUpRequest struct {
	RequestPacket
	Protocol     Protocol
	RaidServerId int64
	IsPractice   bool
}
