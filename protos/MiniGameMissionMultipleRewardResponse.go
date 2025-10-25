package protos

type MiniGameMissionMultipleRewardResponse struct {
	ResponsePacket
	Protocol        Protocol
	AddedHistoryDBs []MissionHistoryDB
	ParcelResultDB  ParcelResultDB
}
