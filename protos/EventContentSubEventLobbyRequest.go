package protos

type EventContentSubEventLobbyRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
