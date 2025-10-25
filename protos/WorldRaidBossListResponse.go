package protos

type WorldRaidBossListResponse struct {
	ResponsePacket
	Protocol        Protocol
	BossListInfoDBs []WorldRaidBossListInfoDB
}
