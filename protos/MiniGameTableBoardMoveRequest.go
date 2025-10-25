package protos

type MiniGameTableBoardMoveRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Steps          []HexLocation
}
