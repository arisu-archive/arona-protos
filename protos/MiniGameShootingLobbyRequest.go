package protos

type MiniGameShootingLobbyRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
