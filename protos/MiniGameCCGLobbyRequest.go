package protos

type MiniGameCCGLobbyRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
