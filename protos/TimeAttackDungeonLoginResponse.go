package protos

type TimeAttackDungeonLoginResponse struct {
	ResponsePacket
	Protocol       Protocol
	PreviousRoomDB TimeAttackDungeonRoomDB
}
