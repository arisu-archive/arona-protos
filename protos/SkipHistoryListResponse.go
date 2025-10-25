package protos

type SkipHistoryListResponse struct {
	ResponsePacket
	Protocol      Protocol
	SkipHistoryDB SkipHistoryDB
}
