package protos

type EventContentDiceRaceRollResponse struct {
	ResponsePacket
	Protocol                  Protocol
	ParcelResultDB            ParcelResultDB
	DiceRaceDB                EventContentDiceRaceDB
	DiceResults               []EventContentDiceResult
	EventContentCollectionDBs []EventContentCollectionDB
}
