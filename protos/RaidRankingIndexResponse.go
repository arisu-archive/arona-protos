package protos

type RaidRankingIndexResponse struct {
	ResponsePacket
	Protocol     Protocol
	RankBrackets []RaidRankBracket
}
