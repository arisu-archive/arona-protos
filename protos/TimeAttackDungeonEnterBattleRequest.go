package protos

type TimeAttackDungeonEnterBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	RoomId        int64
	AssistUseInfo ClanAssistUseInfo
}
