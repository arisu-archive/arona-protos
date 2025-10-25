package protos

type AttendanceRewardResponse struct {
	ResponsePacket
	Protocol              Protocol
	AttendanceBookRewards []AttendanceBookReward
	AttendanceHistoryDBs  []AttendanceHistoryDB
	ParcelResultDB        ParcelResultDB
}
