package protos

type OpenConditionEventListResponse struct {
	ResponsePacket
	Protocol              Protocol
	ConquestTiles         map[int64][]ConquestTileDB
	WorldRaidLocalBossDBs map[int64][]WorldRaidLocalBossDB
}
