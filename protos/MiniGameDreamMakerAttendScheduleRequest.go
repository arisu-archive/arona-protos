package protos

type MiniGameDreamMakerAttendScheduleRequest struct {
	RequestPacket
	Protocol        Protocol
	EventContentId  int64
	ScheduleGroupId int64
}
