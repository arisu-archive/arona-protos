package protos

type MiniGameTableBoardResurrectRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
