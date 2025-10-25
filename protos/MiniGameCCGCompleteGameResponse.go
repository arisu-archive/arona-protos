package protos

type MiniGameCCGCompleteGameResponse struct {
	ResponsePacket
	Protocol       Protocol
	OldSaveDB      MiniGameCCGSaveDB
	ParcelResultDB ParcelResultDB
	RewardParcels  []ParcelInfo
}
