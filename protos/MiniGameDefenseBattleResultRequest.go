package protos

type MiniGameDefenseBattleResultRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageId        int64
	Multiplier     int32
	IsPlayerWin    bool
	BaseDamage     int32
	HeroCount      int32
	AliveCount     int32
	ClearSecond    int32
	Summary        BattleSummary
}
