package protos

type MiniGameTableBoardSyncRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
