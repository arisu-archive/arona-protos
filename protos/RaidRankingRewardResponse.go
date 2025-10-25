package protos

type RaidRankingRewardResponse struct {
	ResponsePacket
	Protocol                Protocol
	ReceivedRankingRewardId int64
	ParcelResultDB          ParcelResultDB
}
