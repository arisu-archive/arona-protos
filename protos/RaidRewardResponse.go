package protos

type RaidRewardResponse struct {
	ResponsePacket
	Protocol         Protocol
	RankingPoint     int64
	BestRankingPoint int64
	ParcelResultDB   ParcelResultDB
}
