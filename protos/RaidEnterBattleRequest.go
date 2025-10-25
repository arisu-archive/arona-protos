package protos

type RaidEnterBattleRequest struct {
	RequestPacket
	Protocol      Protocol
	RaidServerId  int64
	RaidUniqueId  int64
	IsPractice    bool
	EchelonId     int64
	AssistUseInfo ClanAssistUseInfo
}
