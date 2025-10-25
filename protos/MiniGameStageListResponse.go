package protos

type MiniGameStageListResponse struct {
	ResponsePacket
	Protocol           Protocol
	MiniGameHistoryDBs []MiniGameHistoryDB
}
