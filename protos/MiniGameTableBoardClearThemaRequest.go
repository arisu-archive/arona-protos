package protos

type MiniGameTableBoardClearThemaRequest struct {
	RequestPacket
	Protocol                    Protocol
	EventContentId              int64
	PreserveItemEffectUniqueIds []int64
}
