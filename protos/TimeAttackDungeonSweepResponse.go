package protos

type TimeAttackDungeonSweepResponse struct {
	ResponsePacket
	Protocol       Protocol
	Rewards        [][]ParcelInfo
	ParcelResultDB ParcelResultDB
	RoomDB         TimeAttackDungeonRoomDB
}
