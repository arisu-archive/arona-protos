package protos

type EliminateRaidLobbyInfoDB struct {
	RaidLobbyInfoDB
	OpenedBossGroups             []string
	BestRankingPointPerBossGroup map[string]int64
}
