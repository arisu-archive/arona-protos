package protos

type MiniGameTableBoardSweepResponse struct {
	ResponsePacket
	Protocol       Protocol
	SaveDB         TBGBoardSaveDB
	ParcelResultDB ParcelResultDB
}
