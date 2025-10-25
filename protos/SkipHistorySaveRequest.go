package protos

type SkipHistorySaveRequest struct {
	RequestPacket
	Protocol      Protocol
	SkipHistoryDB SkipHistoryDB
}
