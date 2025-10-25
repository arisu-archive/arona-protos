package protos

type MiniGameTableBoardEncounterInputRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	ObjectServerId int64
	EncounterStage int32
	SelectedIndex  int32
}
