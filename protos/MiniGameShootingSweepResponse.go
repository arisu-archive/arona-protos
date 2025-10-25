package protos

type MiniGameShootingSweepResponse struct {
	ResponsePacket
	Protocol       Protocol
	Rewards        [][]ParcelInfo
	ParcelResultDB ParcelResultDB
}
