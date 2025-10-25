package protos

type MiniGameTableBoardSyncResponse struct {
	ResponsePacket
	Protocol Protocol
	SaveDB   TBGBoardSaveDB
}
