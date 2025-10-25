package protos

type WorldRaidLobbyResponse struct {
	ResponsePacket
	Protocol        Protocol
	ClearHistoryDBs []WorldRaidClearHistoryDB
	LocalBossDBs    []WorldRaidLocalBossDB
	BossGroups      []WorldRaidBossGroup
}
