package protos

type EliminateRaidRankingRewardResponse struct {
	ResponsePacket
	Protocol                Protocol
	ReceivedRankingRewardId int64
	ParcelResultDB          ParcelResultDB
}
