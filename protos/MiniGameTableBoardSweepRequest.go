package protos

type MiniGameTableBoardSweepRequest struct {
	RequestPacket
	Protocol                    Protocol
	EventContentId              int64
	PreserveItemEffectUniqueIds []int64
}
