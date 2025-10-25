package protos

type MiniGameMissionListResponse struct {
	ResponsePacket
	Protocol                 Protocol
	MissionHistoryUniqueIds  []int64
	ProgressDBs              []MissionProgressDB
	ClearedOrignalMissionIds []int64
}
