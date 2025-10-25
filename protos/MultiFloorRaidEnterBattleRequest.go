package protos

type MultiFloorRaidEnterBattleRequest struct {
	RequestPacket
	Protocol       Protocol
	SeasonId       int64
	Difficulty     int32
	EchelonId      int32
	AssistUseInfos []ClanAssistUseInfo
}
