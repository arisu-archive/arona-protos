package protos

type BattleSummary struct {
	HashKey            int64
	IsBossBattle       bool
	BattleType         BattleTypes
	StageId            int64
	GroundId           int64
	Winner             GroupTag
	EndType            BattleEndType
	EndFrame           int32
	UnitType           float64
	ResultValue        float64
	Group01Summary     GroupSummary
	Group02Summary     GroupSummary
	WeekDungeonSummary WeekDungeonSummary
	RaidSummary        RaidSummary
	TouchCountSummary  ExcessiveTouch
	ArenaSummary       ArenaSummary
	ContinueCount      int32
	ElapsedRealtime    float32
	IsAbort            bool
	IsDefeatBattle     bool
}
