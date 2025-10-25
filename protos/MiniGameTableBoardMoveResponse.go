package protos

type MiniGameTableBoardMoveResponse struct {
	ResponsePacket
	Protocol       Protocol
	PlayerDB       TBGPlayerDB
	SaveDB         TBGBoardSaveDB
	EncounterDB    TBGEncounterDB
	ParcelResultDB ParcelResultDB
}
