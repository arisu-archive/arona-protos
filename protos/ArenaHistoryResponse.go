package protos

type ArenaHistoryResponse struct {
	ResponsePacket
	Protocol            Protocol
	ArenaHistoryDBs     []ArenaHistoryDB
	ArenaDamageReportDB []ArenaDamageReportDB
}
