package protos

type EliminateRaidCreateBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	RaidUniqueId  int64
	IsPractice    bool
	AssistUseInfo ClanAssistUseInfo
}
