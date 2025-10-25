package protos

type EventContentDiceRaceUseItemResponse struct {
	ResponsePacket
	Protocol                  Protocol
	ParcelResultDB            ParcelResultDB
	DiceRaceDB                EventContentDiceRaceDB
	DiceResults               []EventContentDiceResult
	EventContentCollectionDBs []EventContentCollectionDB
}
