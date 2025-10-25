package protos

type MiniGameCCGSummary struct {
	RandomSeed        int32
	IsPlayerWin       bool
	Strikers          []MiniGameCCGCharacterDB
	TotalUsedCost     int32
	TotalDamageAmount int32
	TotalKillCount    int32
	TotalSkillCount   map[int64]int32
	InputLogs         []string
}
