package protos

type BattlePassMissionSingleRewardResponse struct {
	ResponsePacket
	Protocol       Protocol
	AddedHistoryDB MissionHistoryDB
	ParcelResultDB ParcelResultDB
	BattlePassInfo BattlePassInfoDB
}
