package protos

type MultiFloorRaidEndBattleRequest struct {
	RequestPacket
	Protocol       Protocol
	SeasonId       int64
	Difficulty     int32
	Summary        BattleSummary
	EchelonId      int32
	AssistUseInfos []ClanAssistUseInfo
}
