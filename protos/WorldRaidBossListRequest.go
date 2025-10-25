package protos

type WorldRaidBossListRequest struct {
	RequestPacket
	Protocol                 Protocol
	SeasonId                 int64
	RequestOnlyWorldBossData bool
}
