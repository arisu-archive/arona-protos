package protos

type MiniGameDefenseBattleResultResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	StageHistoryDB MiniGameDefenseStageHistoryDB
}
