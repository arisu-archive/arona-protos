package protos

type MiniGameCCGSweepResponse struct {
	ResponsePacket
	Protocol       Protocol
	Rewards        [][]ParcelInfo
	ParcelResultDB ParcelResultDB
}
