package protos

type MiniGameMissionRewardResponse struct {
	ResponsePacket
	Protocol       Protocol
	AddedHistoryDB MissionHistoryDB
	ParcelResultDB ParcelResultDB
}
