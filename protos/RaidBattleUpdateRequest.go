package protos

type RaidBattleUpdateRequest struct {
	RequestPacket
	Protocol              Protocol
	RaidServerId          int64
	RaidBossIndex         int32
	CumulativeDamage      int64
	CumulativeGroggyPoint int64
}
