package protos

type MomoTalkFavorScheduleResponse struct {
	ResponsePacket
	Protocol             Protocol
	ParcelResultDB       ParcelResultDB
	FavorScheduleRecords map[int64][]int64
}
