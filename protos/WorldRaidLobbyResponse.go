package protos

type WorldRaidLobbyResponse struct {
	ResponsePacket
	ClearHistoryDBs []WorldRaidClearHistoryDB
	LocalBossDBs    []WorldRaidLocalBossDB
	BossGroups      []WorldRaidBossGroup
}
