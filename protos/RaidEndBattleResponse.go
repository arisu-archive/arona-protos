package protos

type RaidEndBattleResponse struct {
	ResponsePacket
	Protocol            Protocol
	RankingPoint        int64
	BestRankingPoint    int64
	ClearTimePoint      int64
	HPPercentScorePoint int64
	DefaultClearPoint   int64
	ParcelResultDB      ParcelResultDB
}
