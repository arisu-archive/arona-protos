package protos

type TimeAttackDungeonEndBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	EchelonId     int32
	RoomId        int64
	Summary       BattleSummary
	AssistUseInfo ClanAssistUseInfo
}
