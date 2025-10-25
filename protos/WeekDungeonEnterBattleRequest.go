package protos

type WeekDungeonEnterBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
	EchelonIndex  int64
}
