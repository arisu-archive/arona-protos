package protos

type MissionRewardResponse struct {
	ResponsePacket
	Protocol       Protocol
	AddedHistoryDB MissionHistoryDB
	ParcelResultDB ParcelResultDB
}
