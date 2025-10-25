package protos

type AttendanceRewardRequest struct {
	RequestPacket
	Protocol               Protocol
	DayByBookUniqueId      map[int64]int64
	AttendanceBookUniqueId int64
	Day                    int64
}
