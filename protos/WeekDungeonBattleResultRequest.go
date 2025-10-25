package protos

type WeekDungeonBattleResultRequest struct {
	RequestPacket
	Protocol           Protocol
	StageUniqueId      int64
	PassCheckCharacter bool
	Summary            BattleSummary
}
