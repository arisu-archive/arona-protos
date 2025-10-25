package protos

type BattlePassMissionListResponse struct {
	ResponsePacket
	Protocol                Protocol
	MissionHistoryUniqueIds []int64
	ProgressDBs             []MissionProgressDB
}
