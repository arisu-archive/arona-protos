package protos

type WorldRaidBossListInfoDB struct {
	GroupId      int64
	WorldBossDB  WorldRaidWorldBossDB
	LocalBossDBs []WorldRaidLocalBossDB
}
