package protos

type MiniGameDefenseGetInfoResponse struct {
	ResponsePacket
	Protocol               Protocol
	EventPointAmount       int64
	DefenseStageHistoryDBs []MiniGameDefenseStageHistoryDB
}
