package protos

type EliminateRaidSweepResponse struct {
	ResponsePacket
	Protocol         Protocol
	TotalSeasonPoint int64
	Rewards          [][]ParcelInfo
	ParcelResultDB   ParcelResultDB
}
