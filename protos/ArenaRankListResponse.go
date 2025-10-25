package protos

type ArenaRankListResponse struct {
	ResponsePacket
	Protocol         Protocol
	TopRankedUserDBs []ArenaUserDB
}
