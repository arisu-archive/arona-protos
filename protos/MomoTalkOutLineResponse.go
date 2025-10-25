package protos

type MomoTalkOutLineResponse struct {
	ResponsePacket
	Protocol             Protocol
	MomoTalkOutLineDBs   []MomoTalkOutLineDB
	FavorScheduleRecords map[int64][]int64
}
