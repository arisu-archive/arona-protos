package protos

type EliminateRaidEndBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	EchelonId     int32
	RaidServerId  int64
	IsPractice    bool
	Summary       BattleSummary
	AssistUseInfo ClanAssistUseInfo
}
