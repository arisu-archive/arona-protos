package protos

type RaidSweepResponse struct {
	ResponsePacket
	Protocol         Protocol
	TotalSeasonPoint int64
	Rewards          [][]ParcelInfo
	ParcelResultDB   ParcelResultDB
}
