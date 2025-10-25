package protos

type WorldRaidLocalBossDB struct {
	SeasonId     int64
	GroupId      int64
	UniqueId     int64
	IsScenario   bool
	IsCleardEver bool
	TacticMscSum int64
	RaidBattleDB RaidBattleDB
	IsContinue   bool
}
