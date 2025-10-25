package protos

type MissionListResponse struct {
	ResponsePacket
	Protocol                 Protocol
	MissionHistoryUniqueIds  []int64
	ProgressDBs              []MissionProgressDB
	DailySuddenMissionInfo   MissionInfo
	ClearedOrignalMissionIds []int64
}
