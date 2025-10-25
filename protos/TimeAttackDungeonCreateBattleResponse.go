package protos

type TimeAttackDungeonCreateBattleResponse struct {
	ResponsePacket
	Protocol       Protocol
	RoomDB         TimeAttackDungeonRoomDB
	ParcelResultDB ParcelResultDB
}
