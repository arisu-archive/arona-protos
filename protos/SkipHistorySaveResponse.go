package protos

type SkipHistorySaveResponse struct {
	ResponsePacket
	Protocol      Protocol
	SkipHistoryDB SkipHistoryDB
}
