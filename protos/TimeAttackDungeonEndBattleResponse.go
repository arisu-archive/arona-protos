package protos

type TimeAttackDungeonEndBattleResponse struct {
	ResponsePacket
	Protocol       Protocol
	RoomDB         TimeAttackDungeonRoomDB
	TotalPoint     int64
	DefaultPoint   int64
	TimePoint      int64
	ParcelResultDB ParcelResultDB
}
