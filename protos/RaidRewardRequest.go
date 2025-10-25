package protos

type RaidRewardRequest struct {
	RequestPacket
	Protocol     Protocol
	RaidServerId int64
	IsPractice   bool
}
