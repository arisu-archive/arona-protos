package protos

type EventContentTreasureLobbyRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
