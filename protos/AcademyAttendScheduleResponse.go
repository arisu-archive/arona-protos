package protos

type AcademyAttendScheduleResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	AcademyDB      AcademyDB
	ExtraRewards   []ParcelInfo
}
