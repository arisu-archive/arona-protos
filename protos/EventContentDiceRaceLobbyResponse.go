package protos

type EventContentDiceRaceLobbyResponse struct {
	ResponsePacket
	Protocol   Protocol
	DiceRaceDB EventContentDiceRaceDB
}
