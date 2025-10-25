package protos

type SchoolDungeonBattleResultRequest struct {
	RequestPacket
	Protocol           Protocol
	StageUniqueId      int64
	PassCheckCharacter bool
	Summary            BattleSummary
}
