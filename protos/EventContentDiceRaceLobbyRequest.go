package protos

type EventContentDiceRaceLobbyRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
