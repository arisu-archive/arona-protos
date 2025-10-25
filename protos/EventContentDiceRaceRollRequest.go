package protos

type EventContentDiceRaceRollRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
