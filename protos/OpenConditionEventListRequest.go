package protos

type OpenConditionEventListRequest struct {
	RequestPacket
	Protocol                   Protocol
	ConquestEventIds           []int64
	WorldRaidSeasonAndGroupIds map[int64]int64
}
