package protos

type AuditGachaStatisticsResponse struct {
	ResponsePacket
	Protocol    Protocol
	GachaResult map[int64]int64
}
