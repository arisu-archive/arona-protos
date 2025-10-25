package protos

type BattlePassMissionMultipleRewardResponse struct {
	ResponsePacket
	Protocol        Protocol
	AddedHistoryDBs []MissionHistoryDB
	ParcelResultDB  ParcelResultDB
	BattlePassInfo  BattlePassInfoDB
}
