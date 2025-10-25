package protos

type WeekDungeonEnterBattleResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	Seed           int32
	Sequence       int32
}
