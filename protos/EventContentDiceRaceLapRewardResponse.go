package protos

type EventContentDiceRaceLapRewardResponse struct {
	ResponsePacket
	Protocol       Protocol
	DiceRaceDB     EventContentDiceRaceDB
	ParcelResultDB ParcelResultDB
}
