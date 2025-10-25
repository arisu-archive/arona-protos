package protos

type EventContentLocationAttendScheduleRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	ZoneId         int64
	Count          int64
}
