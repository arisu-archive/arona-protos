package protos

type MiniGameShootingLobbyResponse struct {
	ResponsePacket
	Protocol   Protocol
	HistoryDBs []MiniGameShootingHistoryDB
}
