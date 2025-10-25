package protos

type MissionMultipleRewardResponse struct {
	ResponsePacket
	Protocol        Protocol
	AddedHistoryDBs []MissionHistoryDB
	ParcelResultDB  ParcelResultDB
}
