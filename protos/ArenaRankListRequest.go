package protos

type ArenaRankListRequest struct {
	RequestPacket
	Protocol   Protocol
	StartIndex int32
	Count      int32
}
