package protos

type MiniGameTableBoardClearThemaResponse struct {
	ResponsePacket
	Protocol       Protocol
	SaveDB         TBGBoardSaveDB
	ParcelResultDB ParcelResultDB
}
