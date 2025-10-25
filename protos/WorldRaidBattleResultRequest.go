package protos

type WorldRaidBattleResultRequest struct {
	RequestPacket
	Protocol       Protocol
	SeasonId       int64
	GroupId        int64
	UniqueId       int64
	EchelonId      int64
	IsPractice     bool
	IsTicket       bool
	Summary        BattleSummary
	AssistUseInfos []ClanAssistUseInfo
}
