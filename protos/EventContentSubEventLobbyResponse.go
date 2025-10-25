package protos

type EventContentSubEventLobbyResponse struct {
	ResponsePacket
	Protocol             Protocol
	EventContentChangeDB EventContentChangeDB
	IsOnSubEvent         bool
}
