package protos

type EliminateRaidRankingIndexResponse struct {
	ResponsePacket
	Protocol     Protocol
	RankBrackets []RaidRankBracket
}
