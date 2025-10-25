package protos

type SchoolDungeonEnterBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
