package protos

type MiniGameTableBoardEncounterInputResponse struct {
	ResponsePacket
	Protocol                  Protocol
	SaveDB                    TBGBoardSaveDB
	EncounterDB               TBGEncounterDB
	PlayerDiceResult          []int32
	PlayerAddDotEffectResult  *int32
	PlayerDicePlayResult      *TBGDiceRollResult
	ParcelResultDB            ParcelResultDB
	EventContentCollectionDBs []EventContentCollectionDB
}
