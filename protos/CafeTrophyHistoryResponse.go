package protos

type CafeTrophyHistoryResponse struct {
	ResponsePacket
	Protocol                    Protocol
	RaidSeasonRankingHistoryDBs []RaidSeasonRankingHistoryDB
}
