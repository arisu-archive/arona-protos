package protos

type MiniGameTableBoardUseItemRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	ItemSlotIndex  int32
	UsedItemId     int64
	IsDiscard      bool
}
