package protos

type MomoTalkFavorScheduleRequest struct {
	RequestPacket
	Protocol   Protocol
	ScheduleId int64
}
